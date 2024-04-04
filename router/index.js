const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require('@apollo/gateway');

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'games', url: 'http://localhost:8080/query' },
        { name: 'players', url: 'http://localhost:8081/query' },
        { name: 'leaderboards', url: 'http://localhost:8082/query' },
    ],
});


const server = new ApolloServer({ gateway, subscriptions: false });

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
