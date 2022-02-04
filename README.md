
# AI-Gopher

A twitter bot written in Go that Uses GPT3 to post a tweet every 24 hours


## Demo

[First post](https://twitter.com/OwenQuinlan7/status/1489522195412185090?s=20&t=hpT6ipcO0Vo5r7_gpP3qLA)

[Second post](https://twitter.com/OwenQuinlan7/status/1489525370143404032?s=20&t=hpT6ipcO0Vo5r7_gpP3qLA)




## Usage

In order to run this you need a [GPT3 api key](https://openai.com/blog/openai-api/) and a [twitter app](https://developer.twitter.com/en/docs/apps/overview) with OAuth 1.0 that has write permissions

Download the executable for your platform and run it in the terminal like this:
```bash
./AI-Gopher[.exe] -k "ConsumerKey" -s "ConsumerSecret" -at "AccessToken" -as "AccessSecret" -g "GPT3ApiKey"
```

As soon as you run it will post it's first tweet and then sleep for 24 hours.

## License

[GPL v2.0](https://choosealicense.com/licenses/gpl-2.0/)


## Acknowledgements

- [go-twitter](https://github.com/dghubble/go-twitter)
- [go-gpt3](https://github.com/PullRequestInc/go-gpt3)
- [oauth1](https://github.com/dghubble/oauth1)
- [go uuid](https://github.com/nu7hatch/gouuid)


## Roadmap

- Additional flags for topic input, sleep time, GPT token limits.

- A better terminal interface (Probabbly using https://charm.sh/)

