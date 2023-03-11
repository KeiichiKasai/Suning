package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/oauth/redirect", Oauth) //第三方登录
	//用户登录
	u := r.Group("/user")
	{
		u.POST("/register", register) //注册
		u.POST("/login", login)       //登录
		u.POST("/change", change)     //修改密码
		u.DELETE("/logout", logout)   //注销
	}
	//个人账户
	a := r.Group("/account")
	{
		a.POST("/setting", setting)        //设置个人信息
		a.GET("/information", information) //查看个人信息
		//钱包
		w := a.Group("/wallet")
		{
			w.POST("/charge", charge) //充值
			w.GET("/view", view)      //查看余额
		}
		//地址页面
		d := a.Group("/address")
		{
			d.POST("/add", addAddress)   //添加地址
			d.DELETE("/del", delAddress) //删除地址
			d.GET("/view", viewAddress)  //查看地址
		}
		//收藏夹
		l := a.Group("/favorite")
		{
			l.POST("/add", like)       //收藏
			l.DELETE("/minus", unlike) //取消收藏
		}
		//订单
		o := a.Group("/order")
		{
			o.GET("/browse", browse)             //查看所有订单
			o.DELETE("/cancel", cancelOrder)     //取消订单
			o.PUT("/confirmOrder", confirmOrder) //确认收货
		}
	}

	//商品
	s := r.Group("/shop")
	{
		//商品页
		s.GET("/allProduct", allProduct) //所有商品
		s.POST("/buyInShop", buyInShop)  //购买

		//购物车
		cart := s.Group("/cart")
		{
			cart.POST("/addInCart", addInCart)     //放入购物车
			cart.DELETE("/delInCart", delInCart)   //删除购物车内商品
			cart.GET("/watchCart", WatchCart)      //查看购物车
			cart.POST("/buyFromCart", buyFromCart) //在购物车内结算商品
		}

		//评论区
		s.POST("/comment", comment)         //评论
		s.POST("/viewComment", viewComment) //查看评论
	}
	r.Run()
}
