#http://makefiletutorial.foofun.cn/
.PHONY:model
go: swag
	goctl api go --api ./api/main.api  -dir .

swag:
	goctl api plugin -plugin goctl-swagger="swagger -filename main.json -host 127.0.0.1:12001  " -api api/main.api -dir ./api
table = ""
hwmodel:
	goctl model mysql datasource -url="dreamebagha:dreamEbagHA-CCa11@tcp(rm-wz976878216hxs1yl0o.mysql.rds.aliyuncs.com:3306)/dreamebaghomework" -table="$(table)" -dir ./model/homework
askmodel:
	goctl model mysql datasource -url="asklearing:asklearing_cKu8ZYlNop@tcp(rm-wz9q0fbp0zx22b3rluo.mysql.rds.aliyuncs.com:3306)/asklearing" -table="$(table)" -dir ./model/askleran
famousmodel:
	goctl model mysql datasource -url="cloudclsrm:CLOUD$_clrm2@tcp(rm-wz976878216hxs1yl0o.mysql.rds.aliyuncs.com:3306)/cloud_classroom" -table="$(table)" -dir ./model/askleran
wismodel:
	goctl model mysql datasource -url="wisdomlesson:5xlCrgggiFlzbRAw@tcp(rm-wz976878216hxs1yl0o.mysql.rds.aliyuncs.com:3306)/dreamebagwisdomlesson" -table="$(table)" -dir ./model/wisdomlesson
prelesson:
	goctl model mysql datasource -url="dreamebagsp:dreamEbagSP-034a1aa@tcp(rm-wz976878216hxs1yl0o.mysql.rds.aliyuncs.com:3306)/dreamebagpreparelesson" -table="$(table)" -dir ./model/preparelesson
lessonServer:
	goctl model mysql datasource -url="dreamebagsa:dreamEbagSA-A34a1aa@tcp(rm-wz976878216hxs1yl0o.mysql.rds.aliyuncs.com:3306)/dreamEbagLessonServerApi" -table="$(table)" -dir ./model/lessonserver
cloudStorage:
	goctl model mysql datasource -url="eloub8gcs22:rb8g*Ct232@tcp(rm-wz976878216hxs1yl0o.mysql.rds.aliyuncs.com:3306)/dreamebagcloudstorage" -table="$(table)" -dir ./model/cloudstorage


struct = "ko"
mongodb:
	goctl model mongo --type $(struct) -dir ./model/burypoint

style:
	goctl api format --dir ./api