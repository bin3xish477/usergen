# usergen
usergen is a basic tool that takes a list of names and generates usernames with common username/email formats.

![](img/usergen.png)

### Usage

```bash
# only usernames, no emails
usergen -n names.txt
# usernames and emails
usergen -n names.txt -d target.com --save usernames.txt
# write generated usernames and emails to stdout
usergen -n names.txt -d target.com
```
