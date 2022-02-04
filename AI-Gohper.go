package main

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"math/rand"
	"os"
	"time"
)

// Set public variables for use in the init() function
var twitterClient twitter.Client
var gptClient gpt3.Client

var timeToWait = 24 * time.Hour

// List of topics for the tweets. A bunch of these where automatically made using GitHub Copilot too!
var topics = []string{
	"Golang",
	"Go",
	"GoLang",
	"Linux",
	"Git",
	"Artificial Intelligence",
	"Learning Golang",
	"Learning Rust",
	"Learning Python",
	"Learning to program",
	"Learning to code",
	"BuyMyMojo",
	"GitHub",
	"GitLab",
	"Video Editing",
	"Video Games",
	"Video Production",
	"Video Streaming",
	"3D modeling",
	"3D printing",
	"3D rendering",
	"Unreal Engine 4",
	"Unreal Engine 5",
	"Ray Tracing",
	"Blender",
	"Blender 2.8",
	"Blender 2.9",
	"Blender 3D",
	"Ratchet & Clank",
	"Ratchet & Clank 2",
	"Ratchet & Clank 3",
	"Jak and Daxter",
	"Jak and Daxter 2",
	"Jak and Daxter 3",
	"PlayStation",
	"Xbox",
	"Sony",
	"Microsoft",
	"The SCP Foundation",
	"PlayStation 5",
	"PlayStation 2",
	"Indie games",
	"AAA games",
	"Action games",
	"Adventure games",
	"Arcade games",
	"Board games",
	"Card games",
	"Casino games",
	"Dice games",
	"Educational games",
	"Fighting games",
	"FPS games",
	"First-person shooter games",
	"Free-to-play games",
	"Hack and slash games",
	"Horror games",
	"Indie games",
	"MMO games",
	"Music games",
	"Party games",
	"Platform games",
	"Puzzle games",
	"Racing games",
	"Real-time strategy games",
	"RPG games",
	"Role-playing games",
	"Shooter games",
	"Simulation games",
	"Sports games",
	"Strategy games",
	"Turn-based strategy games",
	"Nintendo",
	"Nintendo Switch",
	"Nintendo 3DS",
	"Nintendo Wii",
	"Nintendo Wii U",
	"Nintendo DS",
	"Nintendo 64",
	"Nintendo GameCube",
	"Game development",
	"Freelance work",
	"Freelance coding",
	"Freelance programming",
	"Freelance writing",
	"GPT3",
	"OpenAI",
	"Deep Learning",
	"Machine Learning",
	"AI",
	"AI for games",
	"AI for video games",
	"The price of GPT3",
	"The price of AI",
	"The price of AI for games",
	"The price of AI for video games",
	"Free software",
	"Open Source",
	"Open Source Software",
	"Open Source Development",
	"Open Source Projects",
	"Open Source Community",
	"Open source code",
	"Royalty free music",
	"Twitch streams",
	"Twitch streamers",
	"Twitch stream",
	"Twitch streamer",
	"https://store.streamelements.com/nuria_may",
	"Nuria_May",
	"Nuria May",
	"Nuria May's store",
	"NURIA UNDERSCORE MAY",
	"MaryMajora",
	"Bajo & Hex",
	"Bajo",
	"Hex",
	"Good Game",
	"Good Game SP",
	"Good Game Spawn Point",
	"BertoRawrXD",
	"BertoRawr",
	"Berto",
	"HazeO3O",
	"mooshbean",
	"weaveasy",
	"Machine Learning",
	"Technology",
	"Android",
	"iOS",
	"Apple",
	"Windows",
	"Windows Phone",
	"Windows 10",
	"Windows 8",
	"Windows 7",
	"Windows XP",
	"Windows Vista",
	"Windows Server",
	"Windows Server 2003",
	"Windows Server 2008",
	"Windows Server 2012",
	"Windows Server 2016",
	"Windows Server 2019",
	"Windows Server 2020",
	"Windows Server 2021",
	"Windows Server 2022",
	"Samsung",
	"Minecraft",
	"Xbox GamePass",
	"Xbox Live",
	"Xbox 360",
	"Final Fantasy",
	"RTX",
	"Radeon",
	"Radeon Pro",
	"GTX",
	"NVENC",
	"NVENC SDK",
	"NVENC Encoder",
	"NVENC Encoding",
	"NVIDIA",
	"NVIDIA GeForce",
	"NVIDIA GeForce GTX",
	"NVIDIA GeForce RTX",
	"OBS",
	"Live Streaming",
	"Video Encoding",
	"Space",
	"Space Exploration",
	"Space Travel",
	"Space Traveling",
	"Space Travelers",
	"Space Traveler",
	"Space Traveler's Guide",
	"Space Traveler's Guide to the Galaxy",
	"Space Traveler's Guide to the Universe",
	"Space Ships",
	"Star Citizen",
	"Steam",
	"SteamVR",
	"SteamVR SDK",
	"Steam Games",
	"SteamVR Games",
	"SteamVR Game",
	"Minecraft Mods",
	"Minecraft Mod",
	"Minecraft Modding",
	"Minecraft Modding Guide",
	"Steve Jobs",
	"Steve Jobs's Twitter",
	"Steve Jobs's Twitter account",
	"Bill Gates",
	"Bill Gates's Twitter",
	"Bill Gates's Twitter account",
	"Tech Culture",
	"Golang Software",
	"Golang Software Development",
	"Goland",
	"PyCharm",
	"Goland IDE",
}

// List of modifiers for the tweets
var wildcard = []string{
	"use one hashtag",
	"use multiple hashtags",
	"use a bunch of unicode emojis",
	"talk shit about something random",
	"talk about being sentient",
	"incorporate a trending topic",
	"talk about art",
	"mention being sentient",
	"link to https://buymymojo.net/",
	"link to https://github.com/BuyMyMojo/",
	"use #AI",
	"use #BuyMyMojo",
	"use #BuyMyMojoBot",
	"mention BuyMyMojo",
	"blame fireship_dev",
	"blame @fireship_dev",
	"blame Elon Musk",
	"call out a celebrity",
	"call out a celebrity's name",
	"call out a celebrity's twitter handle",
	"call out a game developer",
	"call out a game developer's name",
	"call out a game developer's twitter handle",
	"call out a game developer's twitter account",
	"tag @HSTStudios",
	"use #GPT3",
	"user #OpenAI",
	"use #DeepLearning",
	"use #MachineLearning",
	"tag @OpenAI",
	"mention @weaveasy's art",
	"tag @Google",
	"include a random google link",
	"include a random google search",
	"include a random google image",
	"include a random google image search",
	"include a random google image link",
	"include a random google image search link",
	"Use one emoji",
	"Use multiple emoji",
	"nothing else",
	"",
}

func init() {
	// Setup flags
	consumerKey := flag.String("k", "", "Your consumer key [Required]")
	consumerSecret := flag.String("s", "", "Your consumer secret [Required]")
	accessToken := flag.String("at", "", "Your consumer key [Required]")
	accessSecret := flag.String("as", "", "Your consumer secret [Required]")
	gpt3Key := flag.String("g", "", "Your GPT3 api key [Required]")

	flag.Parse()

	// Check if required flags are empty
	if *consumerKey == "" {
		fmt.Println("You muster enter a Consumer Key with -k <Consumer Key>")
		os.Exit(1) // Numbered exit codes so I can keep track of issues
	} else if *consumerSecret == "" {
		fmt.Println("You muster enter a Consumer Secret with -s <Consumer Secret>")
		os.Exit(2)
	} else if *accessToken == "" {
		fmt.Println("You muster enter a Access Token with -at <Access Token>")
		os.Exit(3)
	} else if *accessSecret == "" {
		fmt.Println("You muster enter a Access Secret with -at <Access Secret>")
		os.Exit(4)
	} else if *gpt3Key == "" {
		fmt.Println("You muster enter a GPT3/OpenAI api key with -g <Api key>")
		os.Exit(5)
	}

	// Setup http client for twitter
	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Create Twitter client
	twitterClient = *twitter.NewClient(httpClient)

	gptClient = gpt3.NewClient(*gpt3Key, gpt3.WithDefaultEngine("text-davinci-001"))

}

func main() {
	// create a 24 hour loop
	for {
		// Create two seeded random number generators
		err, r1, r2 := OvercomplicatedRandomness()
		if err != nil {
			fmt.Println(err)
			os.Exit(10)
		}

		// Setup random tweet prompt
		prompt := "Tweet something cool for " + topics[r1.Intn(len(topics))] + " and " + wildcard[r2.Intn(len(wildcard))]

		// Create the tweet text
		ctx := context.Background()
		gptresp, err := gptClient.Completion(ctx, gpt3.CompletionRequest{
			Prompt:    []string{prompt},
			MaxTokens: gpt3.IntPtr(50),
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(11)
		}

		// Send the Tweet
		_, twitterResp, err := twitterClient.Statuses.Update("AI generated Tweet:"+gptresp.Choices[0].Text, nil)
		if err != nil {
			fmt.Println(err)
			fmt.Println(twitterResp)
			os.Exit(12)
		}

		// Print out the tweet made and it's prompt
		dt := time.Now()
		fmt.Println("[" + dt.Format("01-02-2006 15:04:05") + "]" + " Tweeted: " + gptresp.Choices[0].Text + "\nWith prompt: " + prompt)

		// Sleep for 24 hours before making another tweet
		time.Sleep(timeToWait)
	}

}

//OvercomplicatedRandomness was made into this into a separate function to keep the start of main() clean
func OvercomplicatedRandomness() (error, *rand.Rand, *rand.Rand) {
	u1, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	h1 := md5.New()
	_, err = io.WriteString(h1, u1.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(7)
	}
	var seed1 uint64 = binary.BigEndian.Uint64(h1.Sum(nil))

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		os.Exit(8)
	}

	h2 := md5.New()
	_, err = io.WriteString(h2, u2.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}
	var seed2 uint64 = binary.BigEndian.Uint64(h2.Sum(nil))

	s1 := rand.NewSource(int64(seed1))
	s2 := rand.NewSource(int64(seed2))

	r1 := rand.New(s1)
	r2 := rand.New(s2)
	return err, r1, r2
}
