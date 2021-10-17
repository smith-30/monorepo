## monorepo

CircleCI configuration example with go monorepo.
See **.circleci/config.yml**

## Documentation

- You can run the test for each version of go.
  - You may want to see if it continues to work with older versions
  - Some version wants to start running CI with the latest version
- I am using [cimg/go](https://circleci.com/developer/images/image/cimg/go?utm_source=google&utm_medium=sem&utm_campaign=sem-google-dg--japac-en-dsa-maxConv-auth-brand&utm_term=g_b-_c__dsa_&utm_content=&gclid=CjwKCAjwk6-LBhBZEiwAOUUDp5Pl82KfYV-OGvWY8ZaG-P88sT5lZHQcd7aOOJPrjfJTIeWzyByxIBoCGB0QAvD_BwE) newly provided by circleci.
- gotestsum makes it easier to see where the test failed.
  - [See](https://app.circleci.com/pipelines/github/smith-30/monorepo/21/workflows/26b5eebe-dcab-4aaa-96f6-6ccd37d586af/jobs/43/tests)
- You can easily increase the number of jobs by defining each service.
  - Write Makefile according to ci on the service side.