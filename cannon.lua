function Cannon(parent)
  local c = {}
  c.parent = parent
  c.pos = parent.pos
  c.angle = 0
  c.w = 20
  c.h = 30

  function c:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.rotate(self.angle+self.parent.angle)
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
    lg.pop()
  end

  function c:update(dt)
    local mouseX, mouseY = love.mouse.getPosition()
    local angleToMouse = math.atan2((mouseY - self.pos.y), (mouseX - self.pos.x))
    if (self.angle + 2*math.pi) < (angleToMouse + 2*math.pi) then
      self.angle = self.angle + (TANK_CANNON_ROTATE_RATE*dt)
    else
      self.angle = self.angle - (TANK_CANNON_ROTATE_RATE*dt)
    end
  end

  return c
end
