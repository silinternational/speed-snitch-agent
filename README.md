# Speed Snitch Agent
This repo includes the source for our Speed Snitch agent which we use to monitor and report on internet speeds
at remote locations. The agent communicates with a management API to receive configuration details related to Tasks 
(speed tests, latency tests, etc.). 

This project is in an early prototyping phase though so it is not a fully functional solution yet. There will eventually 
be `speed-snitch-admin-api` and `speed-snitch-admin-ui` projects as well for the central management capabilities. 

## Build Instructions
1. Run `make dep`
2. Build binary (this will create `cmd/speedsnitch/speedsnitch`):
    - Mac: `make mac`
    - Linux: `make linux`
    - Windows: `make windows`

## Release Process
1. Create new release branch for version number from `develop` following convention `release/x.x.x`
2. Update source with new version number
3. Run `make dist`
4. Commit/push changes
5. Create PR from `release/x.x.x` to `develop`
6. After PR reviewed and merged, create PR from `develop` to `master`
7. Review and merge PR to `master`


## Special Instructions for systemd on a Raspberry Pi
On a Raspberry Pi, you can have systemd ensure the speedsnitch executable is always running and
restarts following an update.

1. Copy (as root) the following files into /lib/systemd/system/ ...
    - extras/pi_lib_systemd_system/speedSnitchAgent.service
    - extras/pi_lib_systemd_system/speedSnitchWatcher.service
    - extras/pi_lib_systemd_system/speedSnitchWatcher.path
2. Create symlinks for these files ...
    $ cd /etc/systemd/system
    $ sudo ln -s /lib/systemd/system/speedSnitchAgent.service speedSnitchAgent.service
    $ sudo ln -s /lib/systemd/system/speedSnitchWatcher.service speedSnitchWatcher.service
    $ sudo ln -s /lib/systemd/system/speedSnitchWatcher.path speedSnitchWatcher.path
3. Reload the systemd daemons and start the new services
    $ sudo systemctl daemon-reload
    $ sudo systemctl start speedSnitchWatcher.path
    $ sudo systemctl start speedSnitchWatcher.service
4. In order to view the status ... $ sudo systemctl status speedSnitchAgent.service
5. In order to see more lines from its output ... $ journalctl -e -u speedSnitchAgent.service


## License - MIT
MIT License

Copyright (c) 2018 SIL International

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.