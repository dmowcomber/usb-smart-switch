# usb-power
Allows you to control usb power through a HomeKit device using the [uhubctl](https://github.com/mvp/uhubctl) command and the [hc](https://github.com/brutella/hc) library.
This has only been tested on the Raspberry Pi with Raspberry Pi OS Buster.

## WARNING
The method used to turn off a USB device is a bit heavy handed as it turns off all USB ports and should be used in a headless Raspberry Pi setup where you don't depend a USB keyboard, mouse, or and other USB device.

## Dependencies
`uhubctl` is required and can be installed on debian based distros with apt:
```
sudo apt install -y uhubctl
```

## Run
The `uhubctl` command requires root access so you can build and run with the following command:
```
go build .
sudo ./usb-power
```

## Run at boot
One way to run this command at boot is to add the following to /etc/rc.local:
```
usbPowerDir={path_to_usb-power}
cd ${usbPowerDir}
${usbPowerDir}/usb-power &
```
