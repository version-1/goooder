INSERT INTO materials (display_id, material_type, slug, title) VALUES ($1, $2, $3, $4) RETURNING id;
