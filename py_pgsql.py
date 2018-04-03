import psycopg2
conn = psycopg2.connect(database='', user='', password='',host='', port='')
# 创建数据库
cur = conn.cursor()
cur.execute('''CREATET TABLE COMPANY
	(
		ID 		INT 	PRIMARYKEY 		NOT	NULL,
		NAME	TEXT					NOT NULL,
		AGE		INT 					NOT NULL,
		ADDRESS	CHAR(50),
		SALARY	RAEAL);''')
conn.commit()
conn.close()

# 插入数据
cur.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESSS,SALARY)' VALUE (1, 'Paul', 32, 'California', 20000.00)");
cur.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESSS,SALARY)' VALUE (1, 'Paul', 32, 'California', 20000.00)");
cur.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESSS,SALARY)' VALUE (1, 'Paul', 32, 'California', 20000.00)");
cur.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESSS,SALARY)' VALUE (1, 'Paul', 32, 'California', 20000.00)");
cur.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESSS,SALARY)' VALUE (1, 'Paul', 32, 'California', 20000.00)");
conn.commit()
conn.close()

# 查询操作
cur.execute("SELECT id, name, addresss, salary from COMPANY")
rows = cur.fetchall()
for row in rows:
	print "ID = ", row[0]
	print "NAME = ", row[1]
	print "ADDRESS = ", row[2]
	print "SALARY = ", row[3]
conn.close()

# 更新操作
cur.execute("UPDADE COMPANY set set SALARY = 2500.00 where ID=1")
conn.commit()
