# LINE TOWN election

Welcome to LINE Town. Hi, you must be the new solution engineer TH-HQ sent to
help us. You are in one of the meaningful events in the town. We are about to have a
new mayor! Previously, we handled this event manually, but this time our population
has increased significantly, and also, there are a lot of candidates, as never
happened before. We know some technologies can help, so that's why you come
here to help us figure it out.

After you and LINE Town discussed, You were assigned to build a web application
for citizens (users) to vote for their new mayor. Coming to the first page, it will show
all the candidates. Users can browse and learn about their interested candidates.
And when the time comes, LINE Town's secretariat will open the vote via the new
system. Then, citizens can vote for their choice, a single one and only. Everyone can
see points going to each candidate during the vote in real-time. LINE Town's
secretariat will close the vote following the planned schedule (on their clock) via the
same system. Nobody can submit from now on. Vote's result will be summarized and
replaced on the current vote page. At that moment, all citizens will know who be the
new mayor of LINE town. By having the new system, this chaotic situation will be in
control and make everyone in town happy.

However, this job has three different roles depending on which you were assigned at
first from the TH-HQ. Each one has its own requirements. Therefore, these are the
detailed requirements if you are;

## A solution full-stack engineer

You are the one who mainly focuses on solving the problem without boundaries of
which part you are working on.

1. Implement a page showing the candidate list, voting, and seeing the result.
2. Retrieve and send data through your implemented API.
3. We will allow citizens to vote right on the page when the vote is opened.
4. Implement a voting mechanism, and the vote can be opened or closed by API.
5. Voters can not vote again. We must check their identity by citizen ID both on the
client and the API.
6. The number of votes can be seen in real-time for each candidate on the page.
7. Handle large traffic after opening the vote. You have to make your system
support throughput at least ~100 RPS.
8. When the vote is closed, no one can vote. The current page will calculate the
score from the server-side and sum up who is the new mayor.
9. Unit tests are required both on the front end and the back end.

<!-- ARCHITECTURE -->
## Architecture

The backend architecture for LINE TOWN election project
- Written frontend in Go
- Follow Hexagonal Architecture 

<!-- FOLDER STRUCTURE -->
cmd                         // main applications for this project
└── service                 // main application 
configs                     // configuration file templates or default configs
internal                    // private application and library code.
├── core                    // all the core components (services, models and ports)
│   ├── models              // the domain models
│   ├── ports               // the interfaces definition used to communicate with actors
│   └── services            // the entry points to the core and each one of them implements the corresponding port
├── handlers                // the driver adapters
├── repositories            // the driven adapters
├── routes                  // the connecting path between handlers and endpoints 
mocks                       // mockup data for test
pkg                         // library code that's ok to use by external applications

<!-- DATABASE -->
## Eloquent ORM Relationships
### Entity Database

| Key | Entity        |
| --- | ------------- |
| PK  | id: int       |
|     | enable: bool  |
|     | state: string |

| Key | Candidate         |
| --- | ----------------- |
| PK  | id: int           |
|     | name: string      |
|     | dob: string       |
|     | bioLink: string   |
|     | imageLink: string |
|     | policy: string    |

| Key | Voter              |
| --- | ------------------ |
| PK  | id: int            |
| FK  | candidateId: int   |
|     | nationalId: string |

| Key | Candidate Vote     |
| --- | ------------------ |
| PK  | id: int            |
| FK  | candidateId: int   |
|     | votedCount: string |

### Entity Relationship

| Entity         | Relationships | With Entity    | Foreign Key | Pivot Table |
| :------------- | :------------ | :------------- | :---------- | :---------- |
| Candidate      | Has Many      | Voter          |             |             |
| Candidate      | Has One       | Candidate Vote |             |             |
| Voter          | Belongs To    | Candidate      | candidateId |             |
| Candidate Vote | Belongs To    | Candidate      | candidateId |             |

<!-- ROADMAP -->
## Roadmap

- [*] Setup Project
  - [*] Initial Go modules and setup folder structure
  - [*] Install configure automatic formatting and development helper
  - [*] Install Viper, leveled logging, and crypto encoding modules
  - [*] setup project with sample config & env

- [*] Features/Repositories
  - [*] Initial database and add core modules
  - [*] Create ports and repository

- [ ] Features/Services
  - [ ] Add request and response service model
  - [ ] Create service ports
  - [ ] Add candidate, voter, and election service
  - [ ] Refactor services and repositories
 
- [ ] Features/RESTful API
  - [ ] Initial fiber and router
  - [ ] Add restful endpoint handler for all services
  - [ ] Connect with services and repositories

- [ ] Features/WebSocket and STOMP
  - [ ] Init websocket and key binding
  - [ ] Add socket handler for stream vote service
  - [ ] Refactor code

- [ ] Features/Integration
  - [ ] Add integration test
  - [ ] Complete the system and connect with frontend
  - [ ] Add initial data file in database
  - [ ] Add reset method
  - [ ] Create dockerfile and docker-compose
  - [ ] Refactor code

