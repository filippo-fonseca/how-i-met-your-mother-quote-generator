package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a slice of quotes
var quotes = []struct {
	quote   string
	person  string
}{
	{"Suit up.", "Barney Stinson"},
	{"Sometimes we search for one thing, but discover another.", "Barney Stinson"},
	{"Whenever I'm sad, I stop being sad and be awesome instead.", "Barney Stinson"},
	{"Dude... where's your suit? Just once, when I say 'suit up,' I wish you'd put on a suit.", "Barney Stinson"},
	{"It’s time to let go of the fantasies. It’s time to grow up.", "Barney Stinson"},
	{"That was the night I was born. I rose like a phoenix from her mentholated bosom and strode into the world, Armani-clad and fully awesome.", "Barney Stinson"},
	{"There' s three rules of cheating: 1. It's not cheating if you're not the one who's married. 2. It's not cheating if her name has two adjacent vowels. 3. And it's not cheating if she's from a different area code.", "Barney Stinson"},
	{"It's going to be legen—wait for it—dary.", "Barney Stinson"},
	{"A lie is just a great story that someone ruined with the truth.", "Barney Stinson"},
	{"Think of me as Yoda. Only instead of being little and green, I wear suits and I'm awesome. I'm your bro. I'm Broda.", "Barney Stinson"},
	{"The Bro Code has been around for centuries. Nay… whatever’s more than centuries.", "Barney Stinson"},
	{"It's going to be legen—wait for it, and I hope you're not lactose intolerant, because the second half of the word is—dary! Legendary!", "Barney Stinson"},
	{"ARTICLE 41 A Bro never cries. EXCEPTIONS: Watching Field of Dreams, E.T., or a sports legend retire.", "Barney Stinson"},
	{"I'm also pleased to announce The Bro Code is now available in select airlines across the country. And Lufthansa. 'Der Bro Code' is, like, huge in Germany.", "Barney Stinson"},
	{"You see, kids, right from the moment I met your mom I knew, I have to love this woman as much as I can, and as long as I can, and I can never stop loving her, not even for a second.", "Ted Mosby"},
	{"You can ask the universe for signs all you want, but ultimately we’ll only see what we want to see… when we’re ready to see it.", "Ted Mosby"},
	{"Never underestimate the power of destiny. Because when you least expect it, the littlest thing can cause a ripple effect that changes your life.", "Ted Mosby"},
	{"If you’re not scared, you’re not taking a chance, and if you’re not taking a chance, then what the hell are you doing?", "Ted Mosby"},
	{"I'm sorry, when I'm excited I abbreviate words I shouldn't.", "Ted Mosby"},
	{"Shouldn't we hold out for the person who doesn’t just tolerate our little quirks, but actually kind of likes them?", "Ted Mosby"},
	{"Here's the thing about mistakes: Sometimes, even when you know something’s a mistake, you gotta make it anyway.", "Ted Mosby"},
	{"It's only once you've stopped that you realize how hard it is to start again.", "Ted Mosby"},
	{"The great moments of your life won't necessarily be the things you do, they'll also be the things that happen to you.", "Ted Mosby"},
	{"We’re going to get older, whether we like it or not, so the only question is whether we get on with our lives or desperately cling to the past.", "Ted Mosby"},
	{"All these little parts of the machine constantly working, making sure that you end up exactly where you're supposed to be, exactly when you're supposed to be there.", "Ted Mosby"},
	{"You see, the universe has a plan, kids, and that plan is always in motion.", "Ted Mosby"},
	{"There are two big days in any love story: the day you meet the girl of your dreams and the day you marry her.", "Ted Mosby"},
	{"Here's the secret kids. None of us can vow to be perfect. In the end all we can do is promise to love each other with everything we've got. Because love's the best thing we do.", "Ted Mosby"},
	{"It's a scary thought but it's also kind of wonderful.", "Ted Mosby"},
	{"Love doesn't make sense. I mean, you can't logic your way into or out of it.", "Ted Mosby"},
	{"Every night can't be legendary. If all nights are legendary, no nights are legendary.", "Ted Mosby"},
	{"If you're looking for a word that means caring about someone beyond all rationality and wanting them to have everything they want, no matter how much it destroys you, it's 'love.'", "Ted Mosby"},
	{"Some couples always support each other, and some couples always challenge each other. But is one really better than the other? Yes. Support is better. Way better.", "Ted Mosby"},
	{"Whatever you do in this life, it's not legendary unless your friends are there to see it.", "Ted Mosby"},
	{"And when you love someone, you just don’t stop. Ever.", "Ted Mosby"},
	{"It’s funny: Sometimes you walk into a place and you know you’re exactly where you’re supposed to be.", "Ted Mosby"},
	{"A word of advice: Play along. The more you fight it, the worse it’s gonna get.", "Ted Mosby"},
	{"When you believe in people, people come through.", "Ted Mosby"},
	{"Sometimes our best decisions are the ones that don't make sense at all.", "Ted Mosby"},
	{"And that, kids, is how I met your mother.", "Ted Mosby"},
	{"Oprah wasn't built in a day.", "Lily Aldrin"},
	{"You can’t design your life like a building. It doesn’t work that way. You just have to live it, and it will design itself.", "Lily Aldrin"},
	{"It's just, eventually we're all gonna move on. It's called growing up.", "Lily Aldrin"},
	{"Why say goodbye to the good things?", "Lily Aldrin"},
	{"We struggle so hard to hold on to these things that we know are gonna disappear eventually. And that’s really noble.", "Lily Aldrin"},
	{"The three-day rule is a childish, manipulative mind game. But yeah, you wait three days.", "Lily Aldrin"},
	{"Say goodbye to all the times you felt lost, to all the times it was a 'no' instead of a 'yes,' to all the scrapes and bruises, to all the heartache.", "Lily Aldrin"},
	{"We'll always be friends. It's just never gonna be how it was. It can't be. And that doesn't have to be a sad thing.", "Lily Aldrin"},
	{"You can't just skip ahead to where you think your life should be.", "Lily Aldrin"},
	{"If you keep giving up on people so quickly, you’re gonna miss out on something great.", "Robin Scherbatsky"},
	{"Aren’t you tired of waiting for destiny, Ted? Isn’t it time to make your own destiny?", "Robin Scherbatsky"},
	{"I really don't like feelings.", "Robin Scherbatsky"},
	{"If you have chemistry, all you need is one other thing … Timing.", "Robin Scherbatsky"},
	{"I get recognized one time, and I start thinking I'm Julia Roberts. I’m no VIP; I'm not even an IP; I'm just a lonely little P sitting out here in the gutter.", "Robin Scherbatsky"},
	{"Of course, it’s one thing not to want something; it’s another to be told you can’t have it.", "Robin Scherbatsky"},
	{"Why am I constantly looking for reasons not to be happy?", "Robin Scherbatsky"},
	{"I don’t know where I’m gonna be in five years. I don’t wanna know. I want my life to be an adventure.", "Robin Scherbatsky"},
	{"The future is scary, but you can’t just run back to the past because it's familiar.", "Robin Scherbatsky"},
	{"If I ask you to change too many things about yourself, you’re not gonna be the man I fell in love with.", "Robin Scherbatsky"},
	{"I don’t wanna get married right now, maybe ever, and if we got together I’d feel like I’d either have to marry you or break your heart, and I just couldn’t do either of those things.", "Robin Scherbatsky"},
	{"Oh my god, look at you cowards. So afraid of any kind of change. So terrified of anything new.", "Robin Scherbatsky"},
	{"It's not about proof; it's about faith. Faith is what gives life shape and meaning.", "Marshall Eriksen"},
	{"You have to let me dance my own battles.", "Marshall Eriksen"},
	{"Look at us, riding around in a limo, eating hot dogs… it's like we're the president.", "Marshall Eriksen"},
	{"That's life, you know. We never end up where you think you're gonna be.", "Marshall Eriksen"},
	{"We always get the best advice from the people we love.", "Marshall Eriksen"},
	{"I’m proud of you. And I’m proud of the choices you’ve made. And I’m proud of who you’ve become.", "Marshall Eriksen"},
	{"You have to be a little bit reckless to be brave.", "Marshall Eriksen"},
	{"Sometimes the thing you’re looking for is right in front of you.", "Marshall Eriksen"},
	{"Love is about making the choice to be with someone, even when it’s hard.", "Marshall Eriksen"},
	{"I’m so happy we’re here together.", "Marshall Eriksen"},
}

// RandomQuote returns a random quote from the list
func RandomQuote() string {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	return fmt.Sprintf("'%s' - %s", quote.quote, quote.person)
}