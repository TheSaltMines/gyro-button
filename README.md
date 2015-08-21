## Gryo Button

When you press the button on an Amazon Dash, we increment the Salt
Mines gyro counter.  See the current count [on the
dashboard](http://dashboard.saltmines.us/south).

Inspired by [this article by Ted
Benson](https://medium.com/@edwardbenson/how-i-hacked-amazon-s-5-wifi-button-to-track-baby-data-794214b0bdd8).

## Development and deployment

This runs on the Raspberry Pi in the office.  I develop it on my Mac
and cross compile it for linux/arm.  You will need Go with the
linux/arm toolchain.  If you're on a Mac, this is easy:

    $ brew install --with-cc-common go

Once that's done:

    $ GOOS=linux GOARCH=arm go build

Then scp the `gyro` binary to `pi@dash1.local` and run it:

    $ sudo ./gyro wlan0
