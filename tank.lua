function Tank(id, x, y)
  local t = {}
  t.id = id
  t.x  = x
  t.y  = y
  t.w = 30
  t.h = 50

  function t:draw()
    love.graphics.setColor(1, 1, 1, 1)
    love.graphics.rectangle("line", self.x, self.y, self.w, self.h)
  end

  function t:update(dt)

  end

  table.insert(bodies.players, t)
  return t
end
