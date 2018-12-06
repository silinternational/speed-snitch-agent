#!/usr/bin/env bash

mkdir -p /etc/systemd/system/default.target.wants

ln -s /usr/lib/systemd/system/speedSnitchAgent.service /etc/systemd/system/default.target.wants/speedSnitchAgent.service
ln -s /usr/lib/systemd/system/speedSnitchWatcher.service /etc/systemd/system/default.target.wants/speedSnitchWatcher.service
ln -s /usr/lib/systemd/system/speedSnitchWatcher.path /etc/systemd/system/default.target.wants/speedSnitchWatcher.path

systemctl daemon-reload

systemctl start speedSnitchWatcher.path
systemctl start speedSnitchWatcher.service