# czero import module

**this repository is related with czero module, implemented by c++.**

## czero origin

czero is a cryptocurrency aimed at using cryptography to provide enhanced
privacy for its users compared to other cryptocurrencies such as Bitcoin. 
**czero** is the module which is to generate the functionality to make user
transaction address encrypt which provide strongest protection of user's 
privacy.
because czero source is not open to public yet, so we provide this module for user 
to easy import the module to compile with go-sero project, please see [compile install guide](https://github.com/sero-cash/go-sero/wiki/Installation-Instructions-for-Linux)



### import project

To import the library and command line program, use the following:

	go get -u github.com/sero-cash/go-czero-import
or cd $YOUR_DEV_ENV/src/github.com/sero-cash
    git clone github.com:sero-cash/go-czero-import.git

### Compile

this project put already compiled library in go-czero-import/czero/lib
we will open source the code soon
[https://gitee.com/hyperspace/czero.git]


## License

The go-sero library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-sero binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.