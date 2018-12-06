#!/usr/bin/env bash

rm -rf fpm/opt
rm -rf fpm/usr
rm -rf fpm/etc

mkdir -p fpm/opt/speedsnitch/
mkdir -p fpm/usr/lib/systemd/system
mkdir -p fpm/etc/systemd/system/default.target.wants

# Copy in the systemd speedsnitch files
cp -r /go/src/github.com/silinternational/speed-snitch-agent/extras/pi_lib_systemd_system/* /data/fpm/usr/lib/systemd/system

if [[ "x" == "x$ADMIN_API_BASE_URL" ]]; then
    echo "Missing ADMIN_API_BASE_URL environment variable";
else
    # Set ADMIN_API_BASE_URL key based on environment variable
    sed -i /data/fpm/usr/lib/systemd/system/speedSnitchAgent.service -e "s ADMIN_API_BASE_URL ${ADMIN_API_BASE_URL} "
fi

if [[ "x" == "x$API_KEY" ]]; then
    echo "Missing API_KEY environment variable";
else
    # Set API_KEY key based on environment variable
    sed -i /data/fpm/usr/lib/systemd/system/speedSnitchAgent.service -e "s API_KEY ${API_KEY} "
fi


# Create the speedsnitch binary for linux
cd /go/src/github.com/silinternational/speed-snitch-agent/cmd/speedsnitch && GOOS=linux GOARCH=amd64 go build -o /data/fpm/opt/speedsnitch/speedsnitch

cd /data

fpm -s dir -t deb -n speedsnitch.linux -v ${APP_VERSION} \
  --after-install ./after-install.sh \
  --deb-systemd-restart-after-upgrade \
  ./fpm/opt/speedsnitch/speedsnitch=/opt/speedsnitch/speedsnitch \
  ./fpm/usr/lib/systemd/system/speedSnitchAgent.service=/usr/lib/systemd/system/speedSnitchAgent.service \
  ./fpm/usr/lib/systemd/system/speedSnitchWatcher.service=/usr/lib/systemd/system/speedSnitchWatcher.service \
  ./fpm/usr/lib/systemd/system/speedSnitchWatcher.path=/usr/lib/systemd/system/speedSnitchWatcher.path
