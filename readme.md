# OpenMusic

OpenMusic is a project aimed at creating an open-source platform for music enthusiasts to share and discover new music.

## Features

- Discover new music from various genres
- Share your favorite tracks with the community
- Create and manage playlists
- Follow other users and see their music recommendations

## Installation

To get started with OpenMusic, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/arganaphang/openmusic.git
```

2. Navigate to the project directory:

```bash
cd openmusic
```

3. Install the dependencies:

```bash
just setup
```

4. Run docker compose

```bash
docker compose up -d
```

5. Run database migration

```bash
just migrate-up
```

## Usage

To start the development server, run:

```bash
just dev
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Contact

For any questions or feedback, please reach out to us at [arganaphangquestian@gmail.com](mailto:arganaphangquestian@gmail.com).
