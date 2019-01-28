## Gryo Button

When you press the button on an Amazon Dash, we (used to) increment
the Salt Mines gyro counter.  These days it just restarts the
dashboard GUI when it gets wedged on the Raspberry Pi.

It works by listening for DHCP DISCOVER packets matching the Dash
button's MAC address.

Inspired by [this article by Ted
Benson](https://medium.com/@edwardbenson/how-i-hacked-amazon-s-5-wifi-button-to-track-baby-data-794214b0bdd8).

## Development and deployment

This runs on the Raspberry Pi in the office.  I develop it on my Mac
and cross compile it for linux/arm.  If you're on a Mac, you can
install Go with:

    $ brew install go

Once that's done:

    $ GOOS=linux GOARCH=arm go build

Then scp the `gyro-button` binary to `pi@dash1.local` and copy it to /usr/local/bin:

    $ sudo mv ./gyro-button /usr/local/bin/gyro

And start the service

    $ sudo systemctl restart gyro
