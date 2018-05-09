package model

import (
	"time"
)

type Account struct {
	Name string `json:"name,omitempty"`
	WIF  string `json:"wif,omitempty"`
}

type DNA []byte

func (dna DNA) ID() string {
	return string(dna)
}

type Member struct {
	Name       string    `gorm:"COLUMN:name;PRIMARY_KEY;TYPE:VARCHAR(64);NOT NULL"  json:"name,omitempty"`
	ID         int64     `gorm:"COLUMN:id;NOT NULL;unique" json:"id,omitempty"`
	SigningKey string    `gorm:"COLUMN:signing_key;TYPE:VARCHAR(64);NOT NULL" json:"signing_key,omitempty"`
	CreatedAt  time.Time `gorm:"COLUMN:created_at;" json:"created_at,omitempty"`
}

type Content []byte

type Post struct {
	DNA     string `gorm:"COLUMN:dna;PRIMARY_KEY;TYPE:VARCHAR(255);NOT NULL" json:"dna,omitempty"`
	Author  string `gorm:"COLUMN:author;TYPE:VARCHAR(64);NOT NULL;index:idx_author" json:"author,omitempty"`
	Title   string `gorm:"COLUMN:title;TYPE:VARCHAR(128);NOT NULL" json:"title,omitempty"`
	Content string `gorm:"COLUMN:content;TYPE:TEXT;NOT NULL" json:"content,omitempty"`
	URI     string `gorm:"COLUMN:uri;TYPE:VARCHAR(64);index:idx_author" json:"uri,omitempty"`
	Digest  string `gorm:"COLUMN:digest;TYPE:VARCHAR(64);NOT NULL" json:"digest,omitempty"`
}