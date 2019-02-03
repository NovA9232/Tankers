require "cannon"

function Tank(id, x, y)
  local t = {}
  t.id = id
  t.x  = x
  t.y  = y
  t.angle = 0
  t.vel = Vec2(0, 0)
  t.w = 30
  t.h = 50
  t.cannon = Cannon(t)

  function t:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.translate(self.x+(self.w/2), self.y+(self.h/2))
      lg.rotate(self.angle)
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
    lg.pop()
    self.cannon:draw()
  end

  function t:update(dt)
    if love.keyboard.isDown("w", "up") then
      self.vel.y = self.vel.y - (TANK_ACC*METER)
    end

    if not love.keyboard.isDown("w", "a", "s", "d", "up", "left", "down", "right") then
      self.vel.x = self.vel.x*(TANK_DECEL*math.min(1/dt, 1))
      self.vel.y = self.vel.y*(TANK_DECEL*math.min(1/dt, 1))
    end

    self.x = self.x + self.vel.x*dt
    self.y = self.y + self.vel.y*dt
    self.cannon:update()
  end

  table.insert(entities.players, t)
  return t
end
