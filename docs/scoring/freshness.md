## Freshness Score Implementation

### Tasks Implemented in freshness.go

`1`  Checking for Freshness Keywords in the URL

- If the URL belongs to frequently updated sections (like /news/, /blog/).

- If any of these sections are found, the freshness score is incremented by +30.

`2` Handling Last-Modified Header

- The function retrieves the Last-Modified date from the URL.

- It then interacts with Cassandra (via cassandra.go) to manage and update stored timestamps.


#### Conditioons Handling Last-Modified Header

`(i)` First time fetching a URL, no last modified date in Cassandra, and URL also has no modified date → Return nothing.

`(ii)` If Cassandra has no date but URL has modified date → Return +30 and also update  in Cassandra.

`(iii)` If URL has a modified date that is the same as stored in Cassandra → Do not return anything, do not update Cassandra.

`(iv)` If URL has a modified date that is latest than Cassandra’s stored date → Return +30 and update Cassandra.