package tools

import (
	"itdp-group3-backend/manager"
	"itdp-group3-backend/model/entity"
	"log"
)

// RunSeed : tool for db seeding
func RunSeed(dbc manager.InfraManagerInterface) {
	sqlDB, _ := dbc.DBCon().DB()
	defer sqlDB.Close()

	repoMng := manager.NewRepo(dbc)

	user1 := entity.User{
		Username: "user1",
		Password: "user1",
		Email:    "user1@gmail.com",
		Account: entity.Account{
			RoleID:      0,
			PhoneNumber: "08111",

			BusinessProfile: entity.BusinessProfile{
				CategoryID:   0,
				Address:      "Surabaya",
				ProfileImage: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\img-bp-febd3bfb-ac9d-40d0-b388-f1662cca906e.jpg",
				ProfileBio:   "bio user1",
				GmapsLink:    "maps.google.com",
				DisplayName:  "user1",

				BusinessHours: []entity.BusinessHour{
					entity.BusinessHour{
						Day:       1,
						OpenHour:  "09:00",
						CloseHour: "21.00",
					},
					entity.BusinessHour{
						Day:       2,
						OpenHour:  "09:00",
						CloseHour: "21.00",
					},
					entity.BusinessHour{
						Day:       3,
						OpenHour:  "09:00",
						CloseHour: "21.00",
					},
				},

				BusinessLinks: []entity.BusinessLink{
					entity.BusinessLink{
						Link:  "sopi.co.id",
						Label: "sopi",
					},
					entity.BusinessLink{
						Link:  "lajada.co.id",
						Label: "lajada",
					},
					entity.BusinessLink{
						Link:  "tokopaedi.co.id",
						Label: "tokopaedi",
					},
				},
			},

			Products: []entity.Product{
				entity.Product{
					ProductName:         "meja",
					Price:               250000,
					Description:         "meja minimalis",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\0287a3f2-4aae-44e8-87ea-7e291dd2cf64\\img-product-c472054b-d56d-4710-bd35-e0c7b9ef54e2.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\0287a3f2-4aae-44e8-87ea-7e291dd2cf64\\img-product-6e8be104-e1fb-4c97-b9a0-4c04b4f84f91.png",
				},
				entity.Product{
					ProductName:         "kursi",
					Price:               100000,
					Description:         "kursi minimalis",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\610c371b-b053-4511-a888-764c90a5976f\\img-product-a391a588-1fbb-4404-9972-e30404ed7c2c.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\610c371b-b053-4511-a888-764c90a5976f\\img-product-7b7ca42c-6a5c-4e0a-a967-a874f4c6dfd3.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\610c371b-b053-4511-a888-764c90a5976f\\img-product-9bbdf023-6ac5-43b3-8919-ab5131f63884.png",
				},
				entity.Product{
					ProductName:         "lemari",
					Price:               500000,
					Description:         "lemari minimalis",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\937a4ba2-6909-4989-a925-3819e966421b\\img-product-552caf53-4efa-49d9-b0c6-aa999859326f.png",
				},
				entity.Product{
					ProductName:         "sofa",
					Price:               500000,
					Description:         "sofa minimalis",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\4b84cdeb-6ff5-475b-b957-dea0a2bf0e26\\img-product-96db3c66-1bed-4d58-992b-e0c45b550f23.png",
				},
				entity.Product{
					ProductName:         "cermin",
					Price:               100000,
					Description:         "cermin minimalis",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\b664131d-c7f2-47ba-98b6-5697e8608f79\\img-product-9c200489-83ab-4e99-9312-613124cbbfae.jpg",
				},
				entity.Product{
					ProductName:         "ayunan",
					Price:               100000,
					Description:         "ayunan",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\cc840007-7099-4440-89df-9ba170a44ce6\\img-product-81dd03c9-26d4-479e-adc1-c5c1fa0708fe.png",
				},
				entity.Product{
					ProductName:         "gazebo",
					Price:               100000,
					Description:         "gazebo",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\a850a8bc-9efd-4832-aaab-6d63d0c3987d\\img-product-00ac012d-c68c-4105-94f8-e4e0582d3d48.jpg",
				},
				entity.Product{
					ProductName:         "wastafel",
					Price:               100000,
					Description:         "wastafel",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\7223c1c2-9f32-4cb9-939a-13ebf3d1fa0e\\img-product-01a946f7-ee7d-42fb-8360-f1fa19367f09.jpg",
				},
				entity.Product{
					ProductName:         "pot bunga",
					Price:               100000,
					Description:         "pot bunga",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\afb8f4db-1859-4482-80aa-08b4b035c7b6\\img-product-a90b61e8-8bfe-47b6-8b75-6d4c84cc9c66.png",
				},
				entity.Product{
					ProductName:         "meja makan",
					Price:               100000,
					Description:         "meja makan",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\02f77cdd-a2b1-49ab-ad23-41b829720677\\img-product-e613a8a0-cdd2-4766-b028-0255f71f24e5.jpg",
				},
			},

			Feeds: []entity.Feed{
				entity.Feed{
					CaptionPost: "hai gais ini meja",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\0287a3f2-4aae-44e8-87ea-7e291dd2cf64\\img-product-c472054b-d56d-4710-bd35-e0c7b9ef54e2.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\0287a3f2-4aae-44e8-87ea-7e291dd2cf64\\img-product-6e8be104-e1fb-4c97-b9a0-4c04b4f84f91.png",
				},
				entity.Feed{
					CaptionPost: "hai gais ini kursi",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\610c371b-b053-4511-a888-764c90a5976f\\img-product-a391a588-1fbb-4404-9972-e30404ed7c2c.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\610c371b-b053-4511-a888-764c90a5976f\\img-product-7b7ca42c-6a5c-4e0a-a967-a874f4c6dfd3.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\610c371b-b053-4511-a888-764c90a5976f\\img-product-9bbdf023-6ac5-43b3-8919-ab5131f63884.png",
				},
				entity.Feed{
					CaptionPost: "hai gais ini lemari",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow wow wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\937a4ba2-6909-4989-a925-3819e966421b\\img-product-552caf53-4efa-49d9-b0c6-aa999859326f.png",
				},
				entity.Feed{
					CaptionPost: "hai gais ini sofa",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow wow wow wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\4b84cdeb-6ff5-475b-b957-dea0a2bf0e26\\img-product-96db3c66-1bed-4d58-992b-e0c45b550f23.png",
				},
				entity.Feed{
					CaptionPost: "hai gais ini pot bunga",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow wow wow wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\02f77cdd-a2b1-49ab-ad23-41b829720677\\img-product-e613a8a0-cdd2-4766-b028-0255f71f24e5.jpg",
				},
				entity.Feed{
					CaptionPost: "hai gais ini bukan pot bunga",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow wow wow wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\7223c1c2-9f32-4cb9-939a-13ebf3d1fa0e\\img-product-01a946f7-ee7d-42fb-8360-f1fa19367f09.jpg",
				},
				entity.Feed{
					CaptionPost: "hai gais ini mungkin pot bunga",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "wow wow wow wow bagus bgt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\afb8f4db-1859-4482-80aa-08b4b035c7b6\\img-product-a90b61e8-8bfe-47b6-8b75-6d4c84cc9c66.png",
				},
				entity.Feed{
					CaptionPost: "hai gais ini gajah",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "BUKAN ANJING",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\a850a8bc-9efd-4832-aaab-6d63d0c3987d\\img-product-00ac012d-c68c-4105-94f8-e4e0582d3d48.jpg",
				},
				entity.Feed{
					CaptionPost: "hai gais ini lemari minimalis",
					DetailComments: []entity.DetailComment{
						entity.DetailComment{
							CommentFill: "bagus bgttt",
						},
					},
					DetailMediaFeeds: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\937a4ba2-6909-4989-a925-3819e966421b\\img-product-552caf53-4efa-49d9-b0c6-aa999859326f.png",
				},
			},
		},
	}

	user1.Encode()
	err := repoMng.AuthRepo().CreateUser(&user1)
	if err != nil {
		log.Fatal(err)
	}

	user2 := entity.User{
		Username: "user2",
		Password: "user2",
		Email:    "user2@gmail.com",
		Account: entity.Account{
			RoleID:      0,
			PhoneNumber: "08112",

			BusinessProfile: entity.BusinessProfile{
				CategoryID:   1,
				Address:      "Jakarta",
				ProfileImage: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\img-bp-36a43497-6890-4e6a-96c3-6118e4a929bd.jpg",
				ProfileBio:   "bio user2",
				GmapsLink:    "maps.google.com",
				DisplayName:  "user2",

				BusinessHours: []entity.BusinessHour{
					entity.BusinessHour{
						Day:       1,
						OpenHour:  "10:00",
						CloseHour: "21.00",
					},
					entity.BusinessHour{
						Day:       2,
						OpenHour:  "10:00",
						CloseHour: "21.00",
					},
					entity.BusinessHour{
						Day:       3,
						OpenHour:  "10:00",
						CloseHour: "21.00",
					},
				},

				BusinessLinks: []entity.BusinessLink{
					entity.BusinessLink{
						Link:  "grabfood.co.id",
						Label: "grabfood",
					},
					entity.BusinessLink{
						Link:  "gofood.co.id",
						Label: "gofood",
					},
					entity.BusinessLink{
						Link:  "sopifood.co.id",
						Label: "sopifood",
					},
				},
			},

			Products: []entity.Product{
				entity.Product{
					ProductName:         "nasi goreng",
					Price:               25000,
					Description:         "nasgor enak",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\95a4b6d8-0408-4c87-928a-f2c9a694ce54\\img-product-f7cd83ae-5ade-4958-b567-c875d91c9336.jpg, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\95a4b6d8-0408-4c87-928a-f2c9a694ce54\\img-product-0d9a3fc4-19a9-438a-b5eb-658459697316.png",
				},
				entity.Product{
					ProductName:         "ikan bakar",
					Price:               20000,
					Description:         "ikan bakar enak",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\05f472a7-8c01-44e7-b4b6-79117c8521ed\\img-product-5c587318-44de-4baf-bc35-eadfdcdc400e.jpg",
				},
				entity.Product{
					ProductName:         "es teh",
					Price:               5000,
					Description:         "es teh manis",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\f31b949e-accd-4645-b69a-33c871077b07\\img-product-d81c895b-821b-4ffc-972a-6ee4d6418ac3.jpg",
				},
				entity.Product{
					ProductName:         "es jeruk",
					Price:               5000,
					Description:         "es jeruk",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\cd493052-c219-41ed-a9dd-d7aaf4e6eccf\\img-product-86aa23a7-e9eb-4af1-8192-b26257c45984.jpg",
				},
				entity.Product{
					ProductName:         "nasi bakar",
					Price:               10000,
					Description:         "nasi bakar",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\06980d68-100e-497f-919c-bff9ccc225e2\\img-product-691b521a-e2bf-469f-8f4b-a30fc0a147db.jpg",
				},
				entity.Product{
					ProductName:         "ayam goreng",
					Price:               10000,
					Description:         "ayam goreng",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\1aa849f7-1dd9-4cde-b600-6dbf0c133378\\img-product-fc7abba3-b5bf-45d3-a7e7-c3256705d586.jpg",
				},
				entity.Product{
					ProductName:         "ayam bakar",
					Price:               10000,
					Description:         "ayam bakar",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\309ee4d7-c50d-4dfe-96d9-b7d6bb55eaeb\\img-product-77e21379-0ecc-4863-b9ef-7aea237ff08b.jpg",
				},
				entity.Product{
					ProductName:         "ayam gepuk",
					Price:               10000,
					Description:         "ayam gepuk",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\1821550f-7c07-43c8-b37f-3f96f0cb398d\\img-product-9c6ca83a-d89d-4d44-8a4c-3641a9a81e80.jpg",
				},
				entity.Product{
					ProductName:         "jus alpukat",
					Price:               10000,
					Description:         "jus alpukat",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\06661d34-aef7-4aee-800c-6bf505fe7da0\\img-product-8267329d-e61e-473d-8078-23de764c70a0.jpg",
				},
				entity.Product{
					ProductName:         "jus melon",
					Price:               10000,
					Description:         "jus melom",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\93b9622c-4a3f-4528-9f66-d9350be1089a\\img-product-d4fb150b-f128-4c95-bf3d-9052fee2b324.jpg",
				},
			},

			Feeds: []entity.Feed{
				entity.Feed{
					CaptionPost: "hai gais ini nasgor",
				},
				entity.Feed{
					CaptionPost: "hai gais ini jus melon",
				},
				entity.Feed{
					CaptionPost: "hai gais ini ayam goreng",
				},
				entity.Feed{
					CaptionPost: "hai gais ini ayam bakar",
				},
			},
		},
	}

	user2.Encode()
	err = repoMng.AuthRepo().CreateUser(&user2)
	if err != nil {
		log.Fatal(err)
	}

	user3 := entity.User{
		Username: "user3",
		Password: "user3",
		Email:    "user3@gmail.com",
		Account: entity.Account{
			RoleID:      0,
			PhoneNumber: "08113",

			BusinessProfile: entity.BusinessProfile{
				CategoryID:   1,
				Address:      "Semarang",
				ProfileImage: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\img-bp-df9713be-9e4a-4815-a47c-b9d59c5cb938.jpg",
				ProfileBio:   "bio user3",
				GmapsLink:    "maps.google.com",
				DisplayName:  "user3",

				BusinessHours: []entity.BusinessHour{
					entity.BusinessHour{
						Day:       1,
						OpenHour:  "07:00",
						CloseHour: "21.00",
					},
					entity.BusinessHour{
						Day:       2,
						OpenHour:  "07:00",
						CloseHour: "21.00",
					},
					entity.BusinessHour{
						Day:       3,
						OpenHour:  "07:00",
						CloseHour: "21.00",
					},
				},

				BusinessLinks: []entity.BusinessLink{
					entity.BusinessLink{
						Link:  "grabfood.co.id",
						Label: "grabfood",
					},
					entity.BusinessLink{
						Link:  "gofood.co.id",
						Label: "gofood",
					},
					entity.BusinessLink{
						Link:  "sopifood.co.id",
						Label: "sopifood",
					},
				},
			},

			Products: []entity.Product{
				entity.Product{
					ProductName:         "roti cokelat",
					Price:               25000,
					Description:         "roti cokelat",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\833776b2-eaed-41b1-8e7d-aca41540db5c\\img-product-1586b11f-bb04-421c-a66e-159f7f69476a.png, E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\833776b2-eaed-41b1-8e7d-aca41540db5c\\img-product-a90554ed-b513-48a3-b0e8-2813f16cc57a.png",
				},
				entity.Product{
					ProductName:         "roti strawberry",
					Price:               20000,
					Description:         "roti strawberry",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\891356c8-5024-4477-a80f-4fc3e2c64dd4\\img-product-66c7cb99-b3b7-42da-80e0-3e0006172cc9.png",
				},
				entity.Product{
					ProductName:         "roti keju",
					Price:               5000,
					Description:         "roti keju",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\e9f12314-f3bd-415b-abac-6d4c19579928\\img-product-82fcdc7b-8b54-4dd1-9cc4-be806f13e5f5.jpg",
				},
				entity.Product{
					ProductName:         "roti kacang hijau",
					Price:               5000,
					Description:         "roti kacang hijau",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\3a821dd6-80c1-4833-b052-880b604f5f31\\img-product-c75c3bea-d27a-4170-8cc7-c7542c560524.png",
				},
				entity.Product{
					ProductName:         "croissant",
					Price:               10000,
					Description:         "croissant",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\3b35ceb3-c69d-4f61-812f-c00ad66e0a4f\\img-product-617cbaea-76fc-4912-893b-2c7c05283935.png",
				},
				entity.Product{
					ProductName:         "waffle",
					Price:               10000,
					Description:         "waffle",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\62cc66cf-df62-4bbe-9b34-9d3a5d3b289a\\img-product-3ce1d84a-b773-4857-8467-8748c1c54601.jpg",
				},
				entity.Product{
					ProductName:         "croffle",
					Price:               10000,
					Description:         "croffle",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\3a40a1a7-7092-4650-83c3-314e551b8cbf\\img-product-9b6e39f8-8265-4fdd-a908-644ee8fa9b50.jpg",
				},
				entity.Product{
					ProductName:         "donat",
					Price:               10000,
					Description:         "donat",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\3beb45cc-d74c-47b8-8cc2-75de90373b6c\\img-product-fda528b7-678c-4d7b-90d4-7e4c0809c00c.png",
				},
				entity.Product{
					ProductName:         "bomboloni",
					Price:               10000,
					Description:         "bomboloni",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\24f25a6b-48f2-467e-ac90-dab6f3a4cca9\\img-product-98123dcc-2f4e-45f4-b6c6-f5cf3fffae50.jpg",
				},
				entity.Product{
					ProductName:         "roti sisir",
					Price:               10000,
					Description:         "roti sisir",
					DetailMediaProducts: "E:\\ITDP Sinarmas Mining\\toktok_dev\\img\\products\\4c93e36f-f7da-4adc-9404-50ede17b39b1\\img-product-a3a5492b-d808-4877-b985-a3bffb2eb09c.png",
				},
			},

			Feeds: []entity.Feed{
				entity.Feed{
					CaptionPost: "hai gais ini roti cokelat",
				},
				entity.Feed{
					CaptionPost: "hai gais ini roti keju",
				},
				entity.Feed{
					CaptionPost: "hai gais ini croissant",
				},
				entity.Feed{
					CaptionPost: "hai gais ini waffle",
				},
			},
		},
	}

	user3.Encode()
	err = repoMng.AuthRepo().CreateUser(&user3)
	if err != nil {
		log.Fatal(err)
	}

	// put database seeder(init dummy data) here
}
