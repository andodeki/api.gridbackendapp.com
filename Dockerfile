FROM heroku/heroku:18-build as build

COPY . /app
WORKDIR /app

# Setup buildpack
RUN mkdir -p /tmp/buildpack/gridbackendapp/go /tmp/build_cache /tmp/env
RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

#Execute Buildpack
RUN STACK=heroku-18 /tmp/buildpack/gridbackendapp/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:18

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m gridbackendapp
USER heroku
CMD /app/bin/api.gridbackendapp.com
