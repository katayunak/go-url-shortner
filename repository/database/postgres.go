package database

import "context"

func (p *PostgresDB) GetLongByShort(ctx context.Context, short string) (string, error) {
	long := ""
	err := p.db.DB().WithContext(ctx).
		Model(&URL{}).
		Where("short=?", short).
		Pluck("long", &long).
		Error
	if err != nil {
		return "", err
	}

	return long, nil
}

func (p *PostgresDB) CreateNewURL(ctx context.Context, url *URL) error {
	err := p.db.DB().WithContext(ctx).Create(url).Error
	if err != nil {
		return err
	}
	return nil
}
