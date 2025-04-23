# CMD

## server start

```bash
docker compose up -d
```

## test

```bash
# 도커 시작
docker compose up -d
# 테스트 시작
go test ./test -coverpkg=./internal/... -coverprofile=./tmp/coverage.out
# 결과를 브라우저에서 확인
go tool cover -html=./tmp/coverage.out
```

## swagger 생성

```bash
swag init -g main.go
```

# Dependencies

- GIN
- GORM
- swagger
- dlv

# 폴더구조

- docs : swagger
- internal : 소스코드
  - controller :
  - dto : request의 값을 가져오기 위한 구조체
  - model : db테이블 구조체 및 비지니스로직 구현
  - repository : DB접속
  - service : repository와 model을 이용해 기능 구현
  - routes : 모든 route를 정의
- tools : 공통적으로 사용되는 함수
- scripts : 빌드, 테스트 등 을 위한 스크립트
- test : API를 호출하여 테스트

# API 구조도

![img](./docs/architecture.png)

# Rule

1. interface 사용하지 않음
1. 코멘트 허용
1. import 가능 규칙
   - Controller: Service, DTO, Model
   - Service: Repository, DTO, Model
   - Repository: Model
   - Model: Model (단, Cycling이 되지 않아야한다.)
   - DTO : DTO
1. Controller
   - request 수신 및 response 반환
   - Service를 호출하여 로직 실행
   - 에러 로그는 여기서
1. Service
   - Repository를 호출하여 DB조회
   - Model의 비지니스 로직 실행
   - 트랜잭션 제어
1. Repository
   - GORM으로 DB데이터를 Model에 바인딩
1. Model
   - GORM 태그 설정
   - 비지니스 로직 구현

# TODO

1. response함수 리펙토링
1. transaction
