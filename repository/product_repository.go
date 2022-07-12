package repository

import (
	"golang-mongodb/model"
	"golang-mongodb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
	FindAllProductWithPagination(page, totalDoc int64) ([]model.Product, error)
	Update(id string, by bson.D) error
	Delete(id string) error
	GetById(id string) (model.Product, error)
	GetByCategory(category string) ([]model.Product, error)
}

type productRepository struct {
	db *mongo.Database
}

func (p *productRepository) Add(newProduct *model.Product) error {
	ctx, cancel := utils.InitContext()
	defer cancel()
	newProduct.Id = primitive.NewObjectID()
	_, err := p.db.Collection("products").InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Retrieve() ([]model.Product, error) {
	var products []model.Product
	ctx, cancel := utils.InitContext()
	defer cancel()
	cursor, err := p.db.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) FindAllProductWithPagination(page, totalDoc int64) ([]model.Product, error) {
	var products []model.Product
	ctx, cancel := utils.InitContext()
	defer cancel()

	limit := totalDoc
	skip := limit * (page - 1)

	opts := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := p.db.Collection("products").Find(ctx, bson.M{}, opts)

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) Update(id string, by bson.D) error {
	ctx, cancel := utils.InitContext()
	defer cancel()

	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.db.Collection("products").UpdateOne(ctx, bson.D{{"_id", primitiveId}}, bson.D{{"$set", by}})
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Delete(id string) error {
	ctx, cancel := utils.InitContext()
	defer cancel()

	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.db.Collection("products").DeleteOne(ctx, bson.D{{"_id", bson.D{{"$eq", primitiveId}}}})
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) GetById(id string) (model.Product, error) {
	var product model.Product

	ctx, cancel := utils.InitContext()
	defer cancel()

	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}

	err = p.db.Collection("products").FindOne(ctx, bson.D{{"_id", primitiveId}}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productRepository) GetByCategory(category string) ([]model.Product, error) {
	var products []model.Product

	ctx, cancel := utils.InitContext()
	defer cancel()

	cursor, err := p.db.Collection("products").Find(ctx, bson.D{{"category", category}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
