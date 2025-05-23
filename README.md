
1. Stworzyłem aplikację pogodową w języku Go z wbudowanym serwerem HTTP, obsługą formularza HTML i integracją z OpenWeatherMap API.

2. Stworztłóem plik Dockerfile który używa wieloetapowe budowanie.

3. Stworzyłem workflow GitHub Actions w `.github/workflows/docker.yml`, który:

   - Buduje tymczasowy obraz `linux/amd64` i skanuje go lokalnie przez Trivy
   - W przypadku braku zagrożeń CRITICAL i HIGH buduje i publikuje obraz `linux/amd64, arm64` do GHCR
   - Wykorzystuje cache przez DockerHub (`cache:weather`)

4.  Tajne dane dodałem jako secrets w repozytorium GitHub:

   - `DOCKERHUB_USERNAME`
   - `DOCKERHUB_TOKEN`
   - `GHCR_TOKEN`


