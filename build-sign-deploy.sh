#!/usr/bin/env bash
set -x

# array of target os/arch
targets=( "darwin/amd64" "linux/amd64" "linux/arm" "windows/386" )
distPath="../../dist"

# download gpg keys to use for signing
runny aws s3 cp s3://$KEY_BUCKET/secret.key ./
runny gpg --import secret.key

cd cmd/speedsnitch/
for target in "${targets[@]}"
do
    # Build binary using gox
    gox -osarch="${target}" -output="${distPath}/${target}/speedsnitch"

    # If OS is windows, append .exe to filename before signing
    if [ "${target}" == "windows/386" ]
    then
        fileToSign="${distPath}/${target}/speedsnitch.exe"
    else
        fileToSign="${distPath}/${target}/speedsnitch"
    fi

    # Sign file with GPG
    runny gpg --yes -a -o "${fileToSign}.sig" --detach-sig $fileToSign

done

# Push dist/ to S3 under folder for CI_BRANCH (ex: develop or 1.2.3)
cd ../..
CI_BRANCH=${CI_BRANCH:="unknown"}
aws s3 sync dist/ s3://$DOWNLOAD_BUCKET$DOWNLOAD_BUCKET_PATH/$CI_BRANCH/
