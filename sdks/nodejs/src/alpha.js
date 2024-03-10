// Copyright 2020 Google LLC All Rights Reserved.

//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

const grpc = require('@grpc/grpc-js');

const messages = require('../lib/alpha/alpha_pb');
const servicesPackageDefinition = require('../lib/alpha/alpha_grpc_pb');

const OUT_OF_RANGE = 'OUT_OF_RANGE';

class Alpha {
	constructor(address, credentials) {
		const services = grpc.loadPackageDefinition(servicesPackageDefinition);
		this.client = new services.agones.dev.sdk.alpha.SDK(address, credentials);
	}

	async playerConnect(playerID) {
		const request = new messages.PlayerID();
		request.setPlayerid(playerID);

		return new Promise((resolve, reject) => {
			this.client.playerConnect(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getBool());
				}
			});
		});
	}

	async playerDisconnect(playerID) {
		const request = new messages.PlayerID();
		request.setPlayerid(playerID);

		return new Promise((resolve, reject) => {
			this.client.playerDisconnect(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getBool());
				}
			});
		});
	}

	async setPlayerCapacity(capacity) {
		const request = new messages.Count();
		request.setCount(capacity);

		return new Promise((resolve, reject) => {
			this.client.setPlayerCapacity(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.toObject());
				}
			});
		});
	}

	async getPlayerCapacity() {
		const request = new messages.Empty();

		return new Promise((resolve, reject) => {
			this.client.getPlayerCapacity(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getCount());
				}
			});
		});
	}

	async getPlayerCount() {
		const request = new messages.Empty();

		return new Promise((resolve, reject) => {
			this.client.getPlayerCount(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getCount());
				}
			});
		});
	}

	async isPlayerConnected(playerID) {
		const request = new messages.PlayerID();
		request.setPlayerid(playerID);

		return new Promise((resolve, reject) => {
			this.client.isPlayerConnected(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getBool());
				}
			});
		});
	}

	async getConnectedPlayers() {
		const request = new messages.Empty();

		return new Promise((resolve, reject) => {
			this.client.getConnectedPlayers(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getListList());
				}
			});
		});
	}

	async getCounterCount(key) {
		const request = new messages.GetCounterRequest();
		request.setName(key);
		return new Promise((resolve, reject) => {
			this.client.getCounter(request, (error, response) => {
				if (error) {
					reject(error);
				} else {
					resolve(response.getCount());
				}
			});
		});
	}

	async incrementCounter(key, amount) {
		const request = new messages.CounterUpdateRequest();
		request.setName(key);
		request.setCountdiff(amount);
		return new Promise((resolve, reject) => {
			this.client.updateCounter(request, (error) => {
				if (error) {
					if (error === OUT_OF_RANGE) {
						return resolve(false);
					}
					reject(error);
				} else {
					resolve(true);
				}
			});
		});
	}

	async decrementCounter(key, amount) {
		const request = new messages.CounterUpdateRequest();
		request.setName(key);
		request.setCountdiff(-amount);
		return new Promise((resolve, reject) => {
			this.client.updateCounter(request, (error) => {
				if (error) {
					if (error === OUT_OF_RANGE) {
						return resolve(false);
					}
					reject(error);
				} else {
					resolve(true);
				}
			});
		});
	}
}

module.exports = Alpha;
