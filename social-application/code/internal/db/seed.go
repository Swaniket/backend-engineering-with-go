// Database seed is the process of initializing a new database with some sample data
package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/Swaniket/social/internal/store"
)

var usernames = []string{
	"alice", "bob", "dave", "charlie", "eve",
	"frank", "grace", "heidi", "ivan", "judy",
	"karl", "laura", "mallory", "nancy", "oscar",
	"peggy", "quentin", "rachel", "steve", "trent",
	"ursula", "victor", "wendy", "xavier", "yvonne",
	"zach", "aaron", "bella", "carl", "diana",
	"edward", "fiona", "george", "hannah", "ian",
	"jasmine", "kevin", "linda", "mike", "nina",
	"oliver", "paula", "quincy", "robert", "susan",
	"tom", "uma", "vincent", "will", "zoe",
}

var titles = []string{
	"Morning Routine Hacks",
	"5 Reasons to Start Journaling Today",
	"The Power of Gratitude: How it Can Change Your Life",
	"Why Meditation is Essential for Success",
	"Overcoming Procrastination: Strategies that Work",
	"Building Resilience in a Fast-Paced World",
	"How to Prioritize Your Tasks Effectively",
	"The Benefits of Reading Before Bedtime",
	"Unlocking Creativity with Mindfulness",
	"Developing Self-Discipline for Success",
	"Why Exercise is Good for Mental Health",
	"Time Management Tips for Busy Professionals",
	"The Importance of Setting Boundaries in Relationships",
	"How to Stay Motivated When Working from Home",
	"Effective Communication Skills for Better Relationships",
	"Creating a Morning Routine That Works for You",
	"The Benefits of Taking Breaks Throughout the Day",
	"Why Learning to Say No is Essential for Success",
	"Building Confidence Through Small Wins",
	"Staying Focused in a World of Distractions",
}

var contents = []string{
	"A healthy morning routine involves waking up at least an hour before your first appointment, allowing time to exercise and meditate.",
	"One of the most significant benefits of journaling is that it allows you to reflect on your thoughts and emotions, helping you gain a deeper understanding of yourself.",
	"To cultivate gratitude, try writing down three things you're thankful for each day. This can be as simple as a good cup of coffee or a beautiful sunset.",
	"Meditation has been shown to reduce stress and anxiety by activating the parasympathetic nervous system, promoting relaxation and calmness.",
	"To overcome procrastination, break down large tasks into smaller, manageable chunks, and focus on making progress one step at a time.",
	"Building resilience requires developing coping skills, learning to adapt to changing circumstances, and cultivating a growth mindset.",
	"Effective task prioritization involves understanding your goals and values, identifying the most important tasks, and focusing on those first.",
	"Reading before bedtime can help improve sleep quality by relaxing the mind and reducing stress levels.",
	"Mindfulness practices such as deep breathing and body scan meditation can increase creativity by allowing you to tap into your subconscious mind.",
	"Developing self-discipline involves setting clear goals, creating a schedule, and holding yourself accountable for achieving them.",
	"Regular exercise has been shown to improve mental health by reducing symptoms of anxiety and depression, improving mood, and increasing energy levels.",
	"To manage time effectively, prioritize tasks based on importance and deadlines, avoid multitasking, and take regular breaks to recharge.",
	"Setting boundaries in relationships involves communicating clearly and assertively, establishing healthy limits, and prioritizing self-care.",
	"Staying motivated while working from home requires creating a dedicated workspace, establishing a routine, and minimizing distractions.",
	"Effective communication skills involve active listening, using 'I' statements, and asking open-ended questions to encourage meaningful dialogue.",
	"Creating a morning routine that works for you involves experimenting with different activities, finding what energizes and motivates you, and sticking to it consistently.",
	"Taking regular breaks throughout the day can improve productivity by reducing burnout, increasing focus, and enhancing creativity.",
	"Learning to say no is essential for success because it allows you to prioritize your own needs, maintain healthy boundaries, and avoid overcommitting.",
	"Building confidence involves celebrating small wins, taking calculated risks, and focusing on progress rather than perfection.",
}

var tags = []string{
	"productivity",
	"selfimprovement",
	"mindfulness",
	"motivation",
	"inspiration",
	"personaldevelopment",
	"wellness",
	"selfcare",
	"mentalhealth",
	"anxiety",
	"depression",
	"stressrelief",
	"exercise",
	"nutrition",
	"productivityhacks",
	"goalsetting",
	"time management",
	"leadership",
	"creativity",
	"innovation",
	"careeradvice",
}

var comments = []string{
	"Great advice on prioritizing tasks! I've been trying to implement this into my daily routine and it's made such a big difference.",
	"I love the idea of taking breaks throughout the day. It really helps me stay focused and avoid burnout.",
	"Mindfulness is so important for mental health. Thanks for sharing your insights on how to cultivate mindfulness!",
	"This article was super helpful in understanding the importance of self-care. I'm definitely going to make some changes to prioritize my own needs.",
	"I've been struggling with anxiety lately and this article gave me some really useful tips on managing stress levels.",
	"Love the creativity hacks! They're so simple but really effective.",
	"The advice on setting boundaries was really helpful for me. I've been trying to implement it into my relationships and it's made a huge difference.",
	"Thanks for sharing your expertise on productivity! It's always great to hear from someone who has experience in this area.",
	"I'm definitely going to try out some of the tips from this article. Can't wait to see how they work out!",
}

func Seed(store store.Storage) {
	ctx := context.Background()
	users := generateUsers(100)

	// reset DB
	_, err := store.DB.ExecContext(ctx, `
		TRUNCATE users, posts, comments RESTART IDENTITY CASCADE;
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Insering the users
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(200, users)
	// Insering the post
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", err)
			return
		}
	}

	log.Println("Seeding is completed!")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@email.com",
			Password: "123456",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
