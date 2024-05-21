import axios from 'axios';

const options = {
	method: 'GET',
	url: 'https://linkedin-api8.p.rapidapi.com/',
	params: {
		username: 'username'
	},
	headers: {
		'X-RapidAPI-Key': 'X-RapidAPI-Key',
		'X-RapidAPI-Host': 'linkedin-api8.p.rapidapi.com'
	}
};

try {
	const response = await axios.request(options);
	console.log(response.data);
} catch (error) {
	console.error(error);
}