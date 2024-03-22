/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "shzqe4cj3k39xnc",
    "created": "2024-03-22 08:02:08.378Z",
    "updated": "2024-03-22 08:02:08.378Z",
    "name": "conversations",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "kx6vwm1f",
        "name": "type",
        "type": "select",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "values": [
            "normal",
            "group"
          ]
        }
      },
      {
        "system": false,
        "id": "zmfskxxz",
        "name": "users",
        "type": "relation",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "_pb_users_auth_",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": null,
          "displayFields": null
        }
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("shzqe4cj3k39xnc");

  return dao.deleteCollection(collection);
})
