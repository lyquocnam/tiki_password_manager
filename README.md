### TIKI Password Management
[![Build Status](https://travis-ci.com/lyquocnam/tiki_password_manager.svg?branch=master)](https://travis-ci.com/lyquocnam/tiki_password_manager)

![Screenshot](https://cdn.pbrd.co/images/HXH8p0Q.png)

#### install
- Download [latest version](https://github.com/lyquocnam/tiki_password_manager/releases)
- Run in `cmd` on windows or `terminal` on `mac` and `linux`

#### Document
##### Create user
Create new user with `username` and `password`
```bash
tiki create tiki12 Tik@12 
```

##### Login
Login with `username` and `password`
```bash
tiki login tiki12 Tik@12 
```

##### Validate password
Validate the `password`
```bash
tiki validate_password Tik@12 
```

##### Change password
Change user's `password`
```bash
tiki change_password Tik@12 Tik@19 
```

##### Note
the default username and password is:
```
tiki
Tik@19
```