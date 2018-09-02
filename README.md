# mgo-n-goji

### Pre-requisites

:black_circle: go
:black_circle: mongoDB
:black_circle: node

#### Installing

Clone it and then run

```bash
cd $GOPATH/src && git clone https://github.com/immontilla/mgo-n-goji.git
```
```bash
cd $GOPATH && go install mgo-n-goji/cmd/app/
```

#### Testing

To test, I am going to use Postman and its client Newman. To show the results, I'll use express and newman-reporter-html.

```bash
cd $GOPATH/src/mgo-n-goji/tests/postman/ && npm install && npm run test-n-report
```

#### Running

```bash
$GOPATH/bin/app
```

- Add *bob* as a contact. His mobile is 600000001.

```bash
curl -iX POST "http://localhost:9889/contacts" -H "Content-Type: application/json" -d "{\"nick\":\"bob\",\"mobile\":[\"600000001\"]}" && echo ""
```

- Get *bob*'s contact data

```bash
curl -iX GET "http://localhost:9889/contacts/bob" && echo ""
```

- Get all contacts so far

```bash
curl -iX GET "http://localhost:9889/contacts" && echo ""
```

- Update *bob*'s contact data

```bash
curl -iX PUT "http://localhost:9889/contacts/bob" -H "Content-Type: application/json" -d "{\"nick\":\"bob\",\"mobile\":[\"600000001\"],\"email\":[\"bob@email.us\"]}" && echo ""
```

- Add mobile 600000002 to a *bob*

```bash
curl -iX PATCH "http://localhost:9889/contacts/bob/addMobile/600000002"  && echo ""
```

- Remove mobile 600000001 from *bob*

```bash
curl -iX PATCH "http://localhost:9889/contacts/bob/delMobile/600000001"  && echo ""
```

- Add the email address bob2@email.us to a *bob*

```bash
curl -iX PATCH "http://localhost:9889/contacts/bob/addEmail/bob2@email.us"  && echo ""
```

- Remove the email address bob@email.us from *bob*

```bash
curl -iX PATCH "http://localhost:9889/contacts/bob/delEmail/bob@email.us"  && echo ""
```

- Change *bob* nick to *bobby*

```bash
curl -iX PATCH "http://localhost:9889/contacts/bob/newNick/bobby"  && echo ""
```

- Delete *bobby*

```bash
curl -iX DELETE "http://localhost:9889/contacts/bobby"  && echo ""
```

