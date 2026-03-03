package gpt

const (
	url               = "https://api.openai.com/v1/chat/completions"
	FeedMetaAdsPrompt = "" +
		"you receive input. you need to respond with json format (without comments!) based on the input generate valid description.\n" +
		"json payload you are answering me with:\n\n" +
		"json object { Description: your_data} \n\n" +
		"input: \"%s\"\n"
	model = "gpt-5-nano"
	role  = "user"
)
