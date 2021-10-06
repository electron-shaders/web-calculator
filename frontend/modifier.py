import os

with open("index.html", "r", encoding="utf-8") as f1,open("%s.bak" % "index.html", "w", encoding="utf-8") as f2:
    for line in f1:
        if 'href="' in line:
            line = line.replace('href="', 'href=".')
        if 'src="' in line:
            line = line.replace('src="', 'src=".')
        f2.write(line)
os.remove("index.html")
os.rename("%s.bak" % "index.html", "index.html")
