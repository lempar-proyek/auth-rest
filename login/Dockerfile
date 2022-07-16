# Copyright (C) 2022 Lempar Proyek
# 
# This file is part of Login System.
# 
# Login System is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
# 
# Login System is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU General Public License
# along with Login System.  If not, see <http://www.gnu.org/licenses/>.

FROM golang:1.18-bullseye AS builder

WORKDIR /usr/src/app

RUN go install github.com/revel/cmd/revel@latest

COPY . .

RUN go mod tidy

RUN revel build . dist

FROM golang:1.18-bullseye AS production

WORKDIR /usr/app

COPY --from=builder /usr/src/app/dist/ /usr/app

CMD [ "/usr/app/run.sh" ]