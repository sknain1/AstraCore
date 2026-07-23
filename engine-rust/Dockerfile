FROM rust:1.87-alpine

WORKDIR /app

COPY Cargo.toml Cargo.lock ./

COPY . .

RUN cargo build --release

CMD ["./target/release/engine-rust"]
