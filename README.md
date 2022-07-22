# icom-customer-reports

## running the tool
A binary called `icom-customer-reports` already exists in the root folder. That binary is build for macOS if you have another operating system you might want to build the binary first, to build the binary execute:
```
go build
```

To get more information about running the tool execute:
```
./icom-customer-reports --help
```
The binary has default values for office location and radius as specified and looks for the input file in path `input/customers.txt`. You can specify different parameters for the office location, radius and the input file.

Example usage:
```
./icom-customer-reports -f input/customers.txt
```


## running tests
In the root folder of the project to run tests execute following command:
```
go test -v
```