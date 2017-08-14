package controllers

import (
	"ChiMu/basic"
	"ChiMu/models"
	"ChiMu/services"
	"ChiMu/utils"
)

type ProductController struct {
	basic.BasicController
}

func (c *ProductController) Product() {
	banners := services.GetBannerList()
	r := make(map[string]interface{}, 0)
	r["banners"] = banners
	r["productList"] = services.GetProductList()
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: r}
	c.ServeJSON()
}

func (c *ProductController) ProductDetail() {
	pid, _ := c.GetInt("id")

	product := services.GetProductWithID(pid)
	if utils.IsNil(product) {
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取失败", Status: 0, Data: nil}
		c.ServeJSON()
	}
	// 优惠券列表
	coupons := services.GetCouponByPID(product.ID)
	product.Coupons = coupons
	// 促销列表
	promotions := services.GetPromotionList(product.ID)
	product.Promotions = promotions
	// 评价列表
	comments := services.GetCommentByPID(product.ID)

	for _, v := range comments {
		user := services.GetCommentUserByID(v.UID)
		v.UserName = user.Name
		v.UserAvatar = user.Avatar
	}

	product.Comments = comments
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: product}
	c.ServeJSON()
}

func (c *ProductController) AddProduct() {
	title := c.GetString("title")
	subMessage := c.GetString("submessage")
	price, _ := c.GetFloat("price")
	volume, _ := c.GetInt("volume")
	description := c.GetString("description")
	oriPrice, _ := c.GetFloat("origprice")
	count, _ := c.GetInt("count")
	invalid, _ := c.GetInt("invalid")

	_, fileErr := c.GetFiles("file")
	_, desFileErr := c.GetFiles("descriptionFile")

	if fileErr != nil && desFileErr != nil {
		product := new(models.Product)
		product.Sales = 0
		product.Title = title
		product.SubMessage = subMessage
		product.Price = price
		product.Volume = volume
		product.Description = description
		product.OriginPrice = oriPrice
		product.Count = count
		product.Invalid = invalid

		var images = make([]models.Image, 0)
		var desImages = make([]models.Image, 0)
		imgURLs, _ := c.SaveFiles("file")
		desImgURLs, _ := c.SaveFiles("descriptionFile")

		for _, v := range imgURLs {
			img := new(models.Image)
			img.URL = v
			img.ProductID = product.ID
			img.ProductType = 0
			img.BannerID = 0
			img.CommentID = 0
			images = append(images, *img)
		}

		for _, v := range desImgURLs {
			img := new(models.Image)
			img.URL = v
			img.ProductID = product.ID
			img.ProductType = 1
			img.BannerID = 0
			img.CommentID = 0
			desImages = append(desImages, *img)
		}
		product.Image = images[0].URL
		product.DescriptionImage = desImages[0].URL
		services.AddProduct(product)
		c.Data["json"] = basic.ResInfo{InfoMsg: "添加成功", Status: 1, Data: nil}
		c.ServeJSON()
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加失败", Status: 0, Data: nil}
	c.ServeJSON()
}

func (c *ProductController) ModifyProduct() {
	id, _ := c.GetInt("id")
	title := c.GetString("title")
	subMessage := c.GetString("submessage")
	price, _ := c.GetFloat("price")
	volume, _ := c.GetInt("volume")
	description := c.GetString("description")
	oriPrice, _ := c.GetFloat("origprice")
	count, _ := c.GetInt("count")
	invalid, _ := c.GetInt("invalid")

	_, fileErr := c.GetFiles("file")
	_, desFileErr := c.GetFiles("descriptionFile")

	if fileErr != nil && desFileErr != nil {
		product := services.GetProductWithID(id)
		product.Sales = 0
		product.Title = title
		product.SubMessage = subMessage
		product.Price = price
		product.Volume = volume
		product.Description = description
		product.OriginPrice = oriPrice
		product.Count = count
		product.Invalid = invalid

		var images = make([]models.Image, 0)
		var desImages = make([]models.Image, 0)
		imgURLs, _ := c.SaveFiles("file")
		desImgURLs, _ := c.SaveFiles("descriptionFile")

		for _, v := range imgURLs {
			img := new(models.Image)
			img.URL = v
			img.ProductID = product.ID
			img.ProductType = 0
			img.BannerID = 0
			img.CommentID = 0
			images = append(images, *img)
		}

		for _, v := range desImgURLs {
			img := new(models.Image)
			img.URL = v
			img.ProductID = product.ID
			img.ProductType = 1
			img.BannerID = 0
			img.CommentID = 0
			desImages = append(desImages, *img)
		}
		product.Image = images[0].URL
		product.DescriptionImage = desImages[0].URL
		// FIX ME: - 这里没有删除原有图片...
		services.ModifyProduct(&product)
		c.Data["json"] = basic.ResInfo{InfoMsg: "修改成功", Status: 1, Data: nil}
		c.ServeJSON()
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "修改失败", Status: 0, Data: nil}
	c.ServeJSON()
}
