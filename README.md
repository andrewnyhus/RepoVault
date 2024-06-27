# RepoVault
Provides interface for automated backups of GitHub repositories you care about. Loosely inspired by https://codeberg.org/forgejo/forgejo

TODO:
- [x] Create dockerized go server
- [ ] Create database with table for repository
  - [ ] backup_enabled
  - [ ] owner
  - [ ] name
  - [ ] latest commit date or latest commit sha (TBD)
  - [ ] github clone url
  - [ ] RepoVault clone url
- [ ] Task A: Create task for fetching new repos to back up
  - [x] Authenticate with GitHub API
  - [ ] Fetch account's repos
    - [ ] Insert new records into database, update "last modified" commit or date 
  - [ ] Fetch starred repos
    - [ ] Insert new records into database, update "last modified" commit or date 
- [ ] Task B: Create task for actually backing up the repos
  - [x] Authenticate with GitHub API
  - [ ] Query database for list of repos to back up
  - [ ] Go through each and do a pull if it exists on machine, otherwise do a clone
- [ ] Implement API endpoints
  - [ ] Create an endpoint to kick off task A & B
  - [ ] Create REST endpoints for repository
    - [ ] GET detail
    - [ ] GET list
    - [ ] PATCH should at least be able to update `backup_enabled`, possibly more
- [ ] Create frontend to display info on repos and to toggle their backup status. (Thinking to use svelte)
