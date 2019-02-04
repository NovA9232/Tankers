function Shell(id, x, y, velMag, angle)
  local s = {}
  s.id = id
  s.pos = Vec2(x, y)
  s.w = 2
  s.h = 4
  s.angle = angle
  s.velMag = velMag
  s.vel = Vec2(velMag*math.sin(angle), velMag*math.cos(angle))

  function s:draw()
    lg.push()
      lg.translate(self.pos.x, self.pos.y)
      lg.rotate(-self.angle)

      lg.setColor(1, 0, 0, 1)
      lg.rectangle("fill", -(self.w/2), -(self.h/2), self.w, self.h)
    lg.pop()
  end

  function s:update(dt)
    self.pos.x = self.pos.x + self.vel.x*dt
    self.pos.y = self.pos.y + self.vel.y*dt
  end

  table.insert(entities.proj, s)
  return s
end
