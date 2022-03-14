FROM mongo:nanoserver-1809
COPY . .
RUN main.go