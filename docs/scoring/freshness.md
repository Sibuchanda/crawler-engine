# Freshness Functions Documentation

## FreshnessScore()

**Description:**  
Determines if a URL has fresh content based on timestamps by comparing the last modified time.

**Parameters:**  
- `url` (string) → The URL being checked.  
- `urlLastModified` (time.Time) → The last modified time retrieved from the URL's HTTP header.  
- `previousTime` (time.Time) → The previously stored last modified time from Cassandra.  

**Returns:**  
- `true` → If the `urlLastModified` time is newer than `previousTime`.  
- `false` → If timestamps are missing, equal, or not updated.  

---

## CheckDomain()

**Description:**  
Checks if the given URL belongs to frequently updated sections like news or blogs.

**Parameters:**  
- `url` (string) → The URL we want to check.  

**Returns:**  
- `true` → If the URL contains keywords like `/news/`, `/blog/`, `/latest/`, etc.  
- `false` → If the URL does not belong to any of these sections.  

---

## UpdateLastModified()

**Description:**  
Updates the last modified timestamp of a URL in the Cassandra database if it is newer than the stored timestamp.

**Parameters:**  
- `url` (string) → The URL whose last modified timestamp is being updated.  
- `newTime` (time.Time) → The new last modified timestamp fetched from the HTTP response.  

**Returns:**  
- `nil` → If the update is successful.  
- `error` → If an error occurs while updating the database.  

---

## GetLastModified()

**Description:**  
Fetches the last modified timestamp of a given URL from the Cassandra database.

**Parameters:**  
- `url` (string) → The URL whose last modified timestamp is being retrieved.  

**Returns:**  
- `time.Time` → The stored last modified timestamp.  
- `error` → If the retrieval process fails.
