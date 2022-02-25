## Golang Bindings for Telegraph API

> ⭐️ A star from you, means a lot!

This library represents a convenient wrapper around the Telegra.ph API. This is my first project in golang, so I would love to know my mistakes.

### Features

Here are the most important aspects, shown as pros and cons:

#### Pros

- Auto-generated from API, leaving less space for mistakes.
- Very easy to use.
- **Upload media** to Telegraph, the method which isn't included in the official API.
- Helpful methods like `Prettify()` to pretty-print any object or convert object to JSON.
- In-built `HTML` to `NodeElement` conversion so you can directly use strings with embedded `HTML`
- Highly documented (doc strings) for IDEs like VSCode.

#### Cons

- No custom error handling. Not much panicking. Though official errors will be returned as `error` object. But I guess that's conventional, and maybe convenient?
- The bugs you are going to find.

### Installing

Just install it using the standard `go get` command.

```shell
go get github.com/StarkBotsIndustries/telegraph/v2
```

### Documentation

Docs can be found here : [![Go Reference](https://pkg.go.dev/badge/github.com/StarkBotsIndustries/telegraph.svg)](https://pkg.go.dev/github.com/StarkBotsIndustries/telegraph/v2)

### Example

A project based on this library can be found here : [Telegraph Go Bot](https://github.com/Telegraph-Go-Bot)

### Usage

This is pretty straightforward. First, import the library

```go
import "github.com/StarkBotsIndustries/telegraph/v2"
```

Now you can call any methods. Let's say, **CreateAccount**?

```go
acc, err := telegraph.CreateAccount(
    telegraph.CreateAccountOpts{ShortName: "Go is Love"},
)

acc.AccessToken
>>> abcdefghijklmnopqrstuvwxyz

acc.ShortName
>>> Go is Love
```

Or **CreatePage**

```go
	page, err := telegraph.CreatePage(telegraph.CreatePageOpts{
		Title: "My First Page",
		Content: []telegraph.Node{
			"Hi ",
			telegraph.NodeElement{
				Tag:      "b",
				Children: []telegraph.Node{"Brothers"},
			},
			" and ",
			telegraph.NodeElement{
				Tag:      "code",
				Children: []telegraph.Node{"Sisters"},
			},
		},
		AccessToken: "abcdefghijklmnopqrstuvwxyz",
	})

page.URL
>>> https://telegra.ph/My-First-Page-02-20

page.Path
>>> My-First-Page-02-20
```

You can also directly use HTML using _HTMLContent_ field

```go
page, err := telegraph.CreatePage(telegraph.CreatePageOpts{
    Title: "My First Page",
    HTMLContent: "Hi <b>Brothers</b> and <code>Sisters</code>",
    AccessToken: "abcdefghijklmnopqrstuvwxyz",
})
```

**Pretty Print** an Object / Convert to **JSON**

```go
acc, err := telegraph.CreateAccount(telegraph.CreateAccountOpts{ShortName: "Go is Love"})
prettifiedObject := telegraph.Prettify(acc)

prettifiedObject
>>> {
>>>     "short_name": "Go is Love",
>>>     "author_name": "",
>>>     "author_url": "",
>>>     "access_token": "abcdefghijklmnopqrstuvwxyz",
>>>     "auth_url": "https://edit.telegra.ph/auth/lmnopqrstuvwxyzabcdefghijk",
>>>     "page_count": 0
>>> }
```

**Upload Media** to Telegraph

```go
file, _ := os.Open("file.jpg")
// os.File is io.Reader so just pass it.
link, _ := telegraph.Upload(file, "photo")

link
>>> https://telegra.ph/file/abcdefghijklmnopqrstuvwxyz.jpg
```

Raw **Get** Request

```go
opts := telegraph.CreateAccountOpts{ShortName: "Durov Uncle"}
acc, err := telegraph.Get("createAccount", opts)

acc.AccessToken
>>> abcdefghijklmnopqrstuvwxyz
```

Raw **Post** Request

```go
opts := telegraph.CreateAccountOpts{ShortName: "Durov Uncle"}
acc, err := telegraph.Post("createAccount", opts)

acc.AccessToken
>>> abcdefghijklmnopqrstuvwxyz
```

### Community and Support

Telegram Channel - [StarkBots](https://t.me/StarkBots)

Telegram Chat - [StarkBotsChat](https://t.me/StarkBotsChat)

### Copyright and License

- Copyright (C) 2022 **Stark Bots** <<https://github.com/StarkBotsIndustries>>

- Licensed under the terms of [GNU General Public License v3 or later (GPLv3+)](https://github.com/StarkBotsIndustries/telegraph/blob/master/LICENSE)
