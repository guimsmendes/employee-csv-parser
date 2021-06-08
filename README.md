<img src = "https://i.ibb.co/pyc9J02/download.png">


## Topics
* [Description](#description)
* [Architecture Diagram](#architecture-diagram)
* [Technologies](#technologies)
* [How to Run](#how-to-run)
* [Technical Debts](#technical-debts)
* [Developers](#desenvolvedores)
* [Helpful Links](#helpful-links)


## Description




## Architecture Diagram
The Architecture chosen for this project was the **Clean Architecture**. Following the image below, this project contains:
* An **Entrypoint** Layer, responsible for starting the jobs by reading the Input Csv files on the csv/ folder.
* An **UseCase** Layer, being the Heart of the application, containing the overall business logic.
* A **DataProvider** Layer, responsible for the connection with the Database for persistence and writing the Output files.
* A **Configuration** Layer, which contains the configuration for the whole project, in this case being responsible for the Database connection.
<img src = "https://cdn-media-1.freecodecamp.org/images/YIABVRTHRz58ZiT6W-emBkfNIQUHBelp8t6U">
The project Architecture in details is presented in the image below:
<img src = "architecture-diagram.png">
 

## Technologies

* **go-memdb**: For data persistence, it was used a memory Database in GoLang from **Hashicorp**, found on the link down below:
 	* > https://github.com/hashicorp/go-memdb


## How to Run
### Install Dependencies
* Before

## Technical Debts

## Developers

[<img src="https://i.ibb.co/HF09yK2/IMG-20181112-WA0023.jpg" width=115 > <br> <sub> Guilherme Mendes </sub>](https://github.com/guimsmendes) |
| :---: |  


## Helpful Links
* [Unit testing with GoLang](https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318)
