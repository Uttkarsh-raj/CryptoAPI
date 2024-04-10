# CryptoAPI
<br>
<center>
  An API service for various cryptocurrency operations, including checking holdings of specific coins by different companies (e.g., Ethereum, Bitcoin), converting between different coins, and background data storage of the latest coin information. Currently, the background task for storing data is commented out due to its continuous nature, making it unsuitable for traditional hosting platforms, and instead requiring resources in a private cloud.
</center>
<br><br>

<!--TABLE OF CONTENTS-->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a> 
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a> 
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
  </details>

<!--About the Project-->
  
## About The Project
<center>
  This API service offers convenient endpoints for checking company holdings in popular cryptocurrencies like Etherium and Bitcoin, as well as converting between different coins. It includes a background task that continuously updates and stores the latest coin data in our database. Built with the Gin web framework and MongoDB, our API is robust and scalable, catering to developers and traders seeking reliable cryptocurrency functionalities.
</center>

### Built With
  - **Golang** - An open-source programming language
  - **Gin-Gonic** - A web framework written in Go
<br><br>

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!--GETTING STARTED-->

## Getting Started

To get started with your Golang application, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).

2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.

3. **Initialize Your Project**: Inside your project directory, run the following command to initialize a new Go module:

   ```
   go mod init github.com/your-username/project-name
   ```
   After installing Golang, you can start running your Go project.
4. **Run without Debugging**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to build and execute your Go application:
   ```
   go run main.go
   ```
   This command will compile and execute your Go program without generating a binary file.



### Installation 

Below is an example of how you can instruct your audience on installing and setting up your app.This template doesn't rely on any external dependencies or services.

1. Run this on your terminal (needs docker to be preinstalled):
   ```
   docker run -p 3000:3000 -it uttkarshraj/cryptapi
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Routes

- **Get "/"**
  * Response as :
      - Sucess : True/False
      - Message: Testing endpoint
      
- **Get "/convert"**
  * Request as :
    ```
    {
	  "fromCurrency": "solana",
	  "toCurrency": "dogecoin",
	  "date": "09-04-2024"
    }
    ```
  * Response as :
      - Sucess : True/False
      - Message: "error" : Error
      - Data: Respective data
    ```
    {
    "data": "891.0625074507258 Dogecoin",
    "success": true
    }
    ```
      
- **Get "/companies"**
  * Request as (only etherium/bitcoin are possible values currently):
    ```
    {
	  "currency": "ethereum"
    }
    ```
  * Response as:
      - Sucess : True/False
      - Message: "error" : Error
      - Data: Respective data
    ```
    {
      "data": [
      {
        "name": "Meitu Inc",
        "symbol": "HKG:1357",
        "country": "HK",
        "total_holdings": 31000
      },
      {
        "name": "Mogo Inc.",
        "symbol": "NASDAQ:MOGO",
        "country": "CA",
        "total_holdings": 146
      }
      ],
      "success": true
    }
    ```
  
<!--USAGE EXAMPLES-->

## Usage
1. Conversion of one crypto coin to other.
2. List all the companies and the holdings they have for a coin.
3. Store data to the database and run this task in the background.

## Screenshots:
(For demo pourpose the job is run after every 2 mins in the screenshot.)
<br>
<center>
<img width="1000" src="https://github.com/Uttkarsh-raj/CryptoAPI/assets/106571927/acbbb50f-6865-4c73-baac-fff534fbac4c"></img>
</center>
<br>
<!-- ROADMAP -->

## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [x] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
  - [ ] Hindi
  - [ ] English

  
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--CONTRIBUTING-->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire ,and create.Any contributions you make are *greatly appreciated*.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact
Uttkarsh Raj - https://github.com/Uttkarsh-raj <br>

Project Link: https://github.com/Uttkarsh-raj/CryptoAPI

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
- [Malven's Grid Cheatsheet](https://grid.malven.co/)
- [Img Shields](https://shields.io)
- [GitHub Pages](https://pages.github.com)
- [Font Awesome](https://fontawesome.com)
- [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
