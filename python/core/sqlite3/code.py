import sqlite3

con = sqlite3.connect('local.db')

cur = con.cursor()

cur.execute('''CREATE TABLE IF NOT EXISTS stocks
               (date text, trans text, symbol text, qty real, price real)''')

cur.execute("INSERT INTO stocks VALUES ('2006-01-05','BUY','RHAT',100,35.14)")

con.commit()

con.close()

purchases = [('2006-03-28', 'BUY', 'IBM', 1000, 45.00),
             ('2006-04-05', 'BUY', 'MSFT', 1000, 72.00),
             ('2006-04-06', 'SELL', 'IBM', 500, 53.00),
            ]
cur.executemany('INSERT INTO stocks VALUES (?,?,?,?,?)', purchases)
con.commit()

for row in cur.execute('SELECT * FROM stocks ORDER BY price'):
        print(row)

t = ('RHAT',) # make tuple
print(t)
result = cur.execute('SELECT * FROM stocks WHERE symbol=?', ('RHAT',))
print(result.fetchall())
print(cur.fetchone())    #None, same with result.fetchone()
