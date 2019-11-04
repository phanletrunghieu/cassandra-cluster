# Datacenter > Rack > Server > Node (v)
# Note: It is a best practice to have more than one seed node per datacenter.
# Single datacenter: https://docs.datastax.com/en/ddac/doc/datastax_enterprise/production/DDACsingleDCperWorkloadType.html
# Multi datacenter: https://docs.datastax.com/en/ddac/doc/datastax_enterprise/production/DDACmultiDCperWorkloadType.html
nodetool status
cqlsh -u cassandra -p cassandra
ALTER USER cassandra WITH PASSWORD 'NEW_PASSWORD';
ALTER KEYSPACE system_auth WITH replication = {'class': 'NetworkTopologyStrategy', 'dc1': 1, 'dc2': 1};
CREATE KEYSPACE demo WITH replication = {'class': 'NetworkTopologyStrategy', 'dc1': 1, 'dc2': 1};