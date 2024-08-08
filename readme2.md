### Command [to be updated]

if you are using linux, ensure to give the script execute permissions and run it:
```bash
chmod +x ./googlecli-linux-amd64
```

run the command follow with --help if you need to check what is the valid flag to use:
```bash
./googlecli-linux-amd64 --help
./googlecli-linux-amd64 -e [your@email.com] --help
```

example command to download email with range today until yesterday, category=primary, in folder=inbox
```bash
./googlecli-linux-amd64  getGmailContent --email abdulraufmustakim@gmail.com -d -p --range 0d,1d --category promotions
```

before download the email, make sure you've created folder `output/email` as place to store downloaded email. 


 
