import requests

url = "https://linkedin-api8.p.rapidapi.com/"

querystring = {"username":"username"}

headers = {
	"X-RapidAPI-Key": "X-RapidAPI-Key",
	"X-RapidAPI-Host": "linkedin-api8.p.rapidapi.com"
}

response = requests.get(url, headers=headers, params=querystring)

print(response.json())