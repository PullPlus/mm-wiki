module github.com/chaiyd/mm-wiki

go 1.12

replace github.com/coreos/go-systemd => ./vendor/github.com/coreos/go-systemd
replace github.com/chaiyd/ => ./

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/astaxie/beego v1.12.2
	github.com/fatih/color v1.7.0
	github.com/go-ego/riot v0.0.0-20191215221243-bdbc256e4cbd
	github.com/go-ldap/ldap/v3 v3.1.11
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/shirou/gopsutil v2.19.11+incompatible
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/snail007/go-activerecord v0.0.0-20190813031814-2ac2f3d7cff0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/russross/blackfriday.v2 v2.0.0
)
