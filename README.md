## Gryo Button

![gryo button](https://joeshaw.org/dropbox/gyro-button.jpg)
![gyro truck](http://saltmines.us/wp-content/uploads/2013/10/trailer-300x225.jpg)

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

    $ GOOS=linux GOARCH=arm GOARM=6 go build

Then scp the `gyro` binary to `pi@dash1.local` and copy it to /usr/local/bin:

    $ sudo mv ./gyro /usr/local/bin

And start the service

    $ sudo service gyro start
