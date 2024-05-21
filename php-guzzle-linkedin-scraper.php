<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://linkedin-api8.p.rapidapi.com/?username=username', [
	'headers' => [
		'X-RapidAPI-Host' => 'linkedin-api8.p.rapidapi.com',
		'X-RapidAPI-Key' => 'X-RapidAPI-Key',
	],
]);

echo $response->getBody();