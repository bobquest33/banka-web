# banka-web
Website for internet banking in fictive bank (School project).
Website back-end is written in go programming language. It uses http://www.gorillatoolkit.org/ mux package and sessions in particular.
And front-end uses jquery library for more convenient handling of ajax calls to server.</br>
Author: Matúš Kačmár

Requirements:</br>
1. <b>Local Postgre SQL server with database sovy-bank created with dump file</b> (banka-web/database/dump.sql).</br>
2. <b>GO lang. installed by this tutorial https://golang.org/doc/install<b>

Notes:
- You can change the db credentials (name, user etc.) in banka-web/database/database.go
- In case you find som bugs please create an issue in repository on github.
- Warrning: This code is for educational purposes only.
