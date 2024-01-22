# Playcount monitor

TODO frontend

* map types in user page
* map types filtering in user page
* userpage max width
* mobile compitability
* mapset page
* getMapsetRemainingPendingTime shit
* redux store
* about page
* header search
* ✅ userpage mapset sort
* ✅ header navigation


TODO backend

* tracker getLastTimeTracked on restart
* usercard nginx gzip 
* possible usercard map pagination, separate handler for maps with pagination (page, per_page)
* list most played mapsets for all users with limit offset endpoint
* list users endpoint
* ✅ handle env client id and client secret
* ✅ add worker that would fetch all users from tracking table very 24 hours
* ✅ add total user map plays stats map[string]string to user dto
* ✅ add bigger test dataset
* ✅ add tracking table with users that are being tracked
* ✅ tracking since - created_at отдавать в dto user
* ✅ add tracking list handler and table for tracked users 
