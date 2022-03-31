package mongo_models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//? https://github.com/google/go-github/issues/19
//! как жить дальше:
//* d = ""
//* r := &github.Repository{Description:&d}
//* client.Repositories.Edit("user", "repo", r)

//! доступные тэги для полей структур
//? OmitEmpty  Only include the field if it's not set to the zero value for the type or to
//            empty slices or maps.
//? MinSize    Marshal an integer of a type larger than 32 bits value as an int32, if that's
//            feasible while preserving the numeric value.
//? Truncate   When unmarshaling a BSON double, it is permitted to lose precision to fit within
//            a float32.
//? Inline     Inline the field, which must be a struct or a map, causing all of its fields
//            or keys to be processed as if they were part of the outer struct. For maps,
//            keys must not conflict with the bson keys of other struct fields.
//! как использовать
// type Product struct {
//     ID        primitive.ObjectID `bson:"_id"`
//     ProductId string             `bson:"product_id" json:"product_id"`
//     *SchemalessDocument          `bson:",inline"`
// }
// type SchemalessDocument struct {
//     Others    bson.M             `bson:"others"`
// }
//? Skip       This struct field should be skipped. This is usually denoted by parsing a "-"
//            for the name.

type MCommodity struct {
	Id *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	// может быть несколько наборов данных/линий/графиков на странице
	ChartData            *[]bson.D `bson:"ChartData,omitempty" json:"ChartData,omitempty"`
	CommodityName        *string   `bson:"CommodityName,omitempty" json:"CommodityName,omitempty"`
	PageTitle            *string   `bson:"PageTitle,omitempty" json:"PageTitle,omitempty"`
	CommodityTitle       *string   `bson:"CommodityTitle,omitempty" json:"CommodityTitle,omitempty"`
	CommodityDescription *string   `bson:"CommodityDescription,omitempty" json:"CommodityDescription,omitempty"`
	ChartType            *string   `bson:"ChartType,omitempty" json:"ChartType,omitempty"`
	ChartLineColor       *string   `bson:"ChartLineColor,omitempty" json:"ChartLineColor,omitempty"`
	ChartLineDash        *string   `bson:"ChartLineDash,omitempty" json:"ChartLineDash,omitempty"`
	ChartXScaleLabel     *string   `bson:"ChartXScaleLabel,omitempty" json:"ChartXScaleLabel,omitempty"`
	ChartYScaleLabel     *string   `bson:"ChartYScaleLabel,omitempty" json:"ChartYScaleLabel,omitempty"`
	TooltipLinks         *bson.D   `bson:"TooltipLinks,omitempty" json:"TooltipLinks,omitempty"`
	TooltipOuterLinks    *bson.D   `bson:"TooltipOuterLinks,omitempty" json:"TooltipOuterLinks,omitempty"`
	TooltipInnerLinks    *bson.D   `bson:"TooltipInnerLinks,omitempty" json:"TooltipInnerLinks,omitempty"`
	//! *PageSettings `bson:"inline" json:"inline"`
}

type MUser struct {
	Id          *primitive.ObjectID `bson:"_id,omitempty" json:"_id,ompitempty"`
	Login       *string             `bson:"Login,omitempty" json:"Login,omitempty"`
	Email       *string             `bson:"Email,omitempty" json:"Email,omitempty"`
	Permissions *bson.A             `bson:"Permissions,omitempty" json:"Permissions,omitempty"`
	Password    *string             `bson:"Password,omitempty" json:"Password,omitempty"`
	Salt        *string             `bson:"Salt,omitempty" json:"Salt,omitempty"`
	Status      *UserStatus         `bson:"Status,omitempty" json:"Status,omitempty"`
}

type UserStatus string

const (
	UserStatusDisabled   UserStatus = "disabled"
	UserStatusUnverified UserStatus = "unverified"
	UserStatusActive     UserStatus = "active"
)

type MDivisionSettings struct {
	Id                    *primitive.ObjectID   `bson:"_id,omitempty" json:"_id,ompitempty"`
	MainTitle             *string               `bson:"MainTitle,omitempty" json:"MainTitle,omitempty"`
	RequiredPermission    *string               `bson:"RequiredPermission,omitempty" json:"RequiredPermission,omitempty"`
	HeaderCommodityList   *[]primitive.ObjectID `bson:"HeaderCommodityList,omitempty" json:"HeaderCommodityList,omitempty"`
	MainPageCommodityList *[]primitive.ObjectID `bson:"MainPageCommodityList,omitempty" json:"MainPageCommodityList,omitempty"`
}

type MCommodityTable struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty" json:"_id,ompitempty"`
	TableData *[]TableDivision    `bson:"TableData,omitempty" json:"TableData,ompitempty"`
}

type TableDivision struct {
	DivisionTitle *string `bson:"DivisionTitle,omitempty" json:"DivisionTitle,omitempty"`
	Fields        *bson.D `bson:"Fields,omitempty" json:"Fields,omitempty"`
}

//! существующие коллекции!
// db.createCollection("commodityTables") // "матрицы"
// db.createCollection("commodityCharts") // "графики"
// db.createCollection("divisions")
// db.createCollection("users")
